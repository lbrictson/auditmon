package schema

import (
	"time"

	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

func (Event) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("event_time", "event_name", "username"),
	}
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.Time("event_time").Default(time.Now()),
		field.String("event_name"),
		field.String("username"),
		field.String("resource").Default("-"),
		field.String("source_ip").Default("-"),
		field.String("request_id").Default("-"),
		field.Bool("read_only").Default(false),
		field.JSON("event_data", make(map[string]interface{})),
		field.String("event_source").Default("auditmon_internal"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return nil
}
