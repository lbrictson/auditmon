package storage

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/lbrictson/auditmon/ent/event"

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
	// Make sure the requested directory exists
	os.Mkdir(config.BufferDirectory, os.ModePerm)
	e.client = config.EntClient
	// Start background retry task
	go e.runEventRetryTask()
	return &e
}

func convertEntEventToModelEvent(input ent.Event) models.Event {
	return models.Event{
		EventID:         input.ID.String(),
		EventName:       input.EventName,
		EventTime:       input.EventTime.UTC(),
		Username:        input.Username,
		Resource:        input.Resource,
		EventSource:     input.EventSource,
		SourceIPAddress: input.SourceIP,
		RequestID:       input.RequestID,
		ReadOnly:        input.ReadOnly,
		EventData:       input.EventData,
		UnixTime:        input.EventTime.Unix(),
	}
}

// Create will insert new events into the database, if the database is not available the events will be saved to disk
// and a retry will be attempted later.  An error is only returned if the insert failed as well as writing the file to
// disk.  This could happen because the disk is full or other IO errors
func (s *EventStore) Create(ctx context.Context, events []models.Event) error {
	dbTime, _ := time.LoadLocation("UTC")
	t := time.Time{}
	// Bulk write, much more performant
	bulk := make([]*ent.EventCreate, len(events))
	for i := range events {
		if events[i].EventData == nil {
			events[i].EventData = make(map[string]interface{})
		}
		if events[i].EventTime.String() == t.String() {
			events[i].EventTime = time.Now().In(dbTime)
		} else {
			events[i].EventTime = events[i].EventTime.In(dbTime)
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
			SetSourceIP(events[i].SourceIPAddress).
			SetEventSource(events[i].EventSource)
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

// EventStoreQueryBuilder is used to construct an event query, the only required fields are EndTime, StartTime, Limit and Page
type EventStoreQueryBuilder struct {
	Username    *string
	Resource    *string
	StartTime   time.Time
	EndTime     time.Time
	EventName   *string
	RequestID   *string
	EventSource *string
	EventIP     *string
	ReadOnly    *bool
	Limit       int
	Page        int
}

func (s *EventStore) Query(ctx context.Context, input EventStoreQueryBuilder) ([]models.Event, error) {
	dbTime, _ := time.LoadLocation("UTC")
	var events []models.Event
	q := s.client.Event.Query().Where(event.EventTimeGTE(input.StartTime.In(dbTime))).Where(event.EventTimeLTE(input.EndTime.In(dbTime)))
	if input.Username != nil {
		q = q.Where(event.UsernameEQ(*input.Username))
	}
	if input.Resource != nil {
		q = q.Where(event.ResourceEQ(*input.Resource))
	}
	if input.EventName != nil {
		q = q.Where(event.EventNameEQ(*input.EventName))
	}
	if input.RequestID != nil {
		q = q.Where(event.RequestIDEQ(*input.RequestID))
	}
	if input.EventSource != nil {
		q = q.Where(event.EventSourceEQ(*input.EventSource))
	}
	if input.EventIP != nil {
		q = q.Where(event.SourceIPEQ(*input.EventIP))
	}
	if input.ReadOnly != nil {
		q = q.Where(event.ReadOnlyEQ(*input.ReadOnly))
	}
	entEvents, err := q.Limit(input.Limit).Offset(input.Limit * input.Page).Order(ent.Desc(event.FieldEventTime)).All(ctx)
	if err != nil {
		return events, err
	}
	for _, x := range entEvents {
		events = append(events, convertEntEventToModelEvent(*x))
	}
	sort.Slice(events, func(i, j int) bool {
		return events[i].EventTime.Unix() > events[j].EventTime.Unix()
	})
	return events, nil
}

func (s EventStore) GetByID(ctx context.Context, id string) (models.Event, error) {
	e := models.Event{}
	uu, err := uuid.Parse(id)
	if err != nil {
		return e, errors.New("invalid id")
	}
	entEvent, err := s.client.Event.Get(ctx, uu)
	if err != nil {
		return e, err
	}
	return convertEntEventToModelEvent(*entEvent), nil
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
			err = os.Remove(directory + "/" + x.Name())
			if err != nil {
				log.Errorf("unable to delete event buffer file %v %v", directory+"/"+x.Name(), err)
			}
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
