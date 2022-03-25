package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/lbrictson/auditmon/pkg/models"

	"github.com/lbrictson/auditmon/ent"
)

// fsLock keeps race conditions from occurring when reading/writing files to the same directory
var fsLock sync.Mutex

type EventStore struct {
	client          *ent.Client
	bufferDirectory string
}

type NewEventStoreInput struct {
	EntClient       *ent.Client
	BufferDirectory string // Any trailing / added to the buffer directory is removed
}

// MustNewEventStore expects a configured and live ent database client
// this function will panic if passed a nil database connection
func MustNewEventStore(config NewEventStoreInput) *EventStore {
	e := EventStore{}
	// Remove trailing / on directory names to avoid file saving issues
	if strings.HasSuffix(config.BufferDirectory, "/") {
		config.BufferDirectory = config.BufferDirectory[:len(config.BufferDirectory)-len("/")]
	}
	if config.EntClient == nil {
		panic(fmt.Sprintf("MustNewEventStore requires a database connection"))
	}
	e.bufferDirectory = config.BufferDirectory
	e.client = config.EntClient
	// Start background retry task
	go e.runEventRetryTask()
	return &e
}

// Create will insert new events into the database, if the database is not available the events will be saved to disk
// and a retry will be attempted later.  An error is only returned if the insert failed as well as writing the file to
// disk.  This could happen because the disk is full or other IO errors
func (s *EventStore) Create(ctx context.Context, events []models.Event) error {
	t := time.Time{}
	// Bulk write, much more performant
	bulk := make([]*ent.EventCreate, len(events))
	for i, _ := range events {
		if events[i].EventData == nil {
			events[i].EventData = make(map[string]interface{})
		}
		if events[i].EventTime.String() == t.String() {
			events[i].EventTime = time.Now()
		}
		if events[i].Resource == "" {
			events[i].Resource = "-"
		}
		if events[i].SourceIPAddress == "" {
			events[i].SourceIPAddress = "-"
		}
		if events[i].RequestID == "" {
			events[i].RequestID = "-"
		}
		bulk[i] = s.client.Event.Create().
			SetEventName(events[i].EventName).
			SetEventData(events[i].EventData).
			SetEventTime(events[i].EventTime).
			SetReadOnly(events[i].ReadOnly).
			SetResource(events[i].Resource).
			SetUsername(events[i].Username).
			SetRequestID(events[i].RequestID).
			SetSourceIP(events[i].SourceIPAddress)
	}
	_, err := s.client.Event.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		// Create event in buffer directory for uploading later
		log.Warn("unable to insert events into database, storing in file buffer")
		err = saveEventsToBufferDirectory(events, s.bufferDirectory, uuid.New().String())
		if err != nil {
			return err
		}
	}
	return nil
}

// saveEventsToBufferDirectory will take an array of events and save them into the specified directory with the file
// name being generated as fileID + .json
func saveEventsToBufferDirectory(e []models.Event, directory string, fileID string) error {
	j, err := json.Marshal(&e)
	if err != nil {
		return errors.New(fmt.Sprintf("unable to marshall event to json to store in buffer, events are lost %v", err))
	}
	fsLock.Lock()
	err = ioutil.WriteFile(fmt.Sprintf("%v/%v.json", directory, fileID), j, 0644)
	if err != nil {
		fsLock.Unlock()
		return errors.New(fmt.Sprintf("unable to write event to file in buffer directory, events are lost %v", err))
	}
	fsLock.Unlock()
	return nil
}

// loadSavedBufferEventsFromDirectory will locate all saved events in the specified directory and bundle them into a
// single array, removing the files when complete
func loadSavedBufferEventsFromDirectory(directory string) ([]models.Event, error) {
	fsLock.Lock()
	files, err := os.ReadDir(directory)
	if err != nil {
		fsLock.Unlock()
		return nil, err
	}
	var events []models.Event
	for _, x := range files {
		// Only pick up json event files
		if strings.Contains(x.Name(), ".json") {
			b, err := ioutil.ReadFile(directory + "/" + x.Name())
			if err != nil {
				fsLock.Unlock()
				return nil, err
			}
			var fileEvents []models.Event
			err = json.Unmarshal(b, &fileEvents)
			if err != nil {
				fsLock.Unlock()
				return nil, err
			}
			events = append(events, fileEvents...)
			os.Remove(directory + "/" + x.Name())
		}
	}
	fsLock.Unlock()
	return events, nil
}

// runEventRetryTask is an endless loop that attempts to insert failed events, on each loop it will compact ones that
// failed to be inserted into a single file for bulk insertion in the future.  Events are retried every 5 minutes
func (s EventStore) runEventRetryTask() {
	for {
		time.Sleep(1 * time.Minute)
		// Gather all failed events
		events, err := loadSavedBufferEventsFromDirectory(s.bufferDirectory)
		if err != nil {
			log.Errorf("unable to load failed events from file system for retry %v", err)
			continue
		}
		// Only proceed if there are events to deal with, most of the time there won't be any
		if events != nil {
			if len(events) != 0 {
				err = s.Create(context.Background(), events)
				if err != nil {
					log.Errorf("something went wrong trying to bulk insert retry events %v", err)
					continue
				}
			}
		}
	}
}
