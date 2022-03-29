package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// UsernameAutofill holds the schema definition for the UsernameAutofill entity.
type UsernameAutofill struct {
	ent.Schema
}

// Fields of the UsernameAutofill.
func (UsernameAutofill) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Unique(),
	}
}

// Edges of the UsernameAutofill.
func (UsernameAutofill) Edges() []ent.Edge {
	return nil
}
