package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EventNameAutofill holds the schema definition for the EventNameAutofill entity.
type EventNameAutofill struct {
	ent.Schema
}

// Fields of the EventNameAutofill.
func (EventNameAutofill) Fields() []ent.Field {
	return []ent.Field{
		field.String("eventName").Unique(),
	}
}

// Edges of the EventNameAutofill.
func (EventNameAutofill) Edges() []ent.Edge {
	return nil
}
