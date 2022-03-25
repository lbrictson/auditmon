package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			StorageKey("oid"),
		field.String("username").Unique(),
		field.String("hashed_password").NotEmpty(),
		field.String("role").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.Time("password_last_set_at").Default(time.Now()),
		field.Bool("force_password_change").Default(false),
		field.Time("last_login").Default(time.Now()),
		field.Int("failed_logins").Default(0),
		field.Bool("locked").Default(false),
		field.Time("locked_until").Default(time.Now()),
		field.String("mfa_secret").Optional(),
		field.Strings("recent_passwords").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
