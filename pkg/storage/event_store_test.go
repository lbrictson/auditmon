package storage

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"

	"github.com/google/uuid"

	"github.com/lbrictson/auditmon/ent"
	"github.com/lbrictson/auditmon/pkg/models"
)

func createEntInMemoryDatabaseClient() *ent.Client {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		panic(err)
	}
	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}

func TestEventStore_Create(t *testing.T) {
	type fields struct {
		client          *ent.Client
		bufferDirectory string
	}
	type args struct {
		ctx   context.Context
		event []models.Event
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "happy path new event store create",
			fields: fields{
				client:          createEntInMemoryDatabaseClient(),
				bufferDirectory: "test_data/",
			},
			args: args{
				ctx: context.Background(),
				event: []models.Event{
					models.Event{
						EventName:       "TestEvent",
						EventTime:       time.Now(),
						Username:        "automated_testing",
						Resource:        "auditmon_test_suite",
						EventSource:     "auditmon",
						SourceIPAddress: "127.0.0.1",
						RequestID:       "aaaa-ffff-eeee-dddd-gggg",
						ReadOnly:        true,
						EventData: map[string]interface{}{
							"purpose": "testing",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "happy path new event store create with nil event data",
			fields: fields{
				client:          createEntInMemoryDatabaseClient(),
				bufferDirectory: "test_data/",
			},
			args: args{
				ctx: context.Background(),
				event: []models.Event{
					models.Event{
						EventName:       "TestEvent",
						EventTime:       time.Now(),
						Username:        "automated_testing",
						Resource:        "auditmon_test_suite",
						EventSource:     "auditmon",
						SourceIPAddress: "127.0.0.1",
						RequestID:       "aaaa-ffff-eeee-dddd-gggg",
						ReadOnly:        true,
						EventData:       nil,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EventStore{
				client:          tt.fields.client,
				bufferDirectory: tt.fields.bufferDirectory,
			}
			if err := s.Create(tt.args.ctx, tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMustNewEventStore(t *testing.T) {
	client := createEntInMemoryDatabaseClient()
	s := EventStore{
		client:          client,
		bufferDirectory: "test_data", // validate that trailing / is trimmed
	}
	type args struct {
		config NewEventStoreInput
	}
	tests := []struct {
		name string
		args args
		want EventStore
	}{
		{
			name: "Create event store happy path",
			args: args{config: NewEventStoreInput{
				EntClient:       client,
				BufferDirectory: "test_data/",
			}},
			want: s,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MustNewEventStore(tt.args.config); !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("MustNewEventStore() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_saveEventToBufferDirectory(t *testing.T) {
	uu := uuid.New().String()
	type args struct {
		e         []models.Event
		directory string
		fileID    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "happy path save event to disk",
			args: args{e: []models.Event{
				models.Event{
					EventID:         uu,
					EventName:       "test_event",
					EventTime:       time.Now(),
					Username:        "auditmon_test",
					Resource:        "auditmon_test",
					EventSource:     "auditmon_test",
					SourceIPAddress: "127.0.0.1",
					RequestID:       "aaa-eee-fff-qqq",
					ReadOnly:        false,
					EventData: map[string]interface{}{
						"test": "data",
					},
				},
			},
				directory: "../../test_data",
				fileID:    uu,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveEventsToBufferDirectory(tt.args.e, tt.args.directory, tt.args.fileID); (err != nil) != tt.wantErr {
				t.Errorf("saveEventToBufferDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Clean up tmp files
			os.Remove(fmt.Sprintf("%v/%v.json", tt.args.directory, tt.args.fileID))
		})
	}
}

func Test_loadSavedBufferEventsFromDirectory(t *testing.T) {
	type args struct {
		directory string
	}
	eventSample := []models.Event{
		{
			EventName: "sample1",
			Username:  "sample1",
		},
		{
			EventName: "sample2",
			Username:  "sample2",
		},
		{
			EventName: "sample3",
			Username:  "sample3",
		},
	}
	err := saveEventsToBufferDirectory(eventSample, "../../test_data", uuid.New().String())
	if err != nil {
		t.Errorf("error saving events to directory to setup for read from directory %v", err)
		return
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Event
		wantErr bool
	}{
		{
			name:    "happy path read events stored in buffer back into memory",
			args:    args{directory: "../../test_data"},
			want:    eventSample,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := loadSavedBufferEventsFromDirectory(tt.args.directory)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadSavedBufferEventsFromDirectory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadSavedBufferEventsFromDirectory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
