// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldHashedPassword holds the string denoting the hashed_password field in the database.
	FieldHashedPassword = "hashed_password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldPasswordLastSetAt holds the string denoting the password_last_set_at field in the database.
	FieldPasswordLastSetAt = "password_last_set_at"
	// FieldForcePasswordChange holds the string denoting the force_password_change field in the database.
	FieldForcePasswordChange = "force_password_change"
	// FieldLastLogin holds the string denoting the last_login field in the database.
	FieldLastLogin = "last_login"
	// FieldFailedLogins holds the string denoting the failed_logins field in the database.
	FieldFailedLogins = "failed_logins"
	// FieldLocked holds the string denoting the locked field in the database.
	FieldLocked = "locked"
	// FieldLockedUntil holds the string denoting the locked_until field in the database.
	FieldLockedUntil = "locked_until"
	// FieldMfaSecret holds the string denoting the mfa_secret field in the database.
	FieldMfaSecret = "mfa_secret"
	// FieldMfaSetupCompleted holds the string denoting the mfa_setup_completed field in the database.
	FieldMfaSetupCompleted = "mfa_setup_completed"
	// FieldRecentPasswords holds the string denoting the recent_passwords field in the database.
	FieldRecentPasswords = "recent_passwords"
	// FieldMfaImage holds the string denoting the mfa_image field in the database.
	FieldMfaImage = "mfa_image"
	// FieldTimezone holds the string denoting the timezone field in the database.
	FieldTimezone = "timezone"
	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldHashedPassword,
	FieldRole,
	FieldCreatedAt,
	FieldPasswordLastSetAt,
	FieldForcePasswordChange,
	FieldLastLogin,
	FieldFailedLogins,
	FieldLocked,
	FieldLockedUntil,
	FieldMfaSecret,
	FieldMfaSetupCompleted,
	FieldRecentPasswords,
	FieldMfaImage,
	FieldTimezone,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// HashedPasswordValidator is a validator for the "hashed_password" field. It is called by the builders before save.
	HashedPasswordValidator func(string) error
	// RoleValidator is a validator for the "role" field. It is called by the builders before save.
	RoleValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt time.Time
	// DefaultPasswordLastSetAt holds the default value on creation for the "password_last_set_at" field.
	DefaultPasswordLastSetAt time.Time
	// DefaultForcePasswordChange holds the default value on creation for the "force_password_change" field.
	DefaultForcePasswordChange bool
	// DefaultLastLogin holds the default value on creation for the "last_login" field.
	DefaultLastLogin time.Time
	// DefaultFailedLogins holds the default value on creation for the "failed_logins" field.
	DefaultFailedLogins int
	// DefaultLocked holds the default value on creation for the "locked" field.
	DefaultLocked bool
	// DefaultLockedUntil holds the default value on creation for the "locked_until" field.
	DefaultLockedUntil time.Time
	// DefaultMfaSetupCompleted holds the default value on creation for the "mfa_setup_completed" field.
	DefaultMfaSetupCompleted bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
