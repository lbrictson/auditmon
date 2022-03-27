package models

import "time"

type User struct {
	ID                    string    `json:"id,omitempty"`
	Username              string    `json:"username,omitempty"`
	HashedPassword        string    `json:"-"`
	Role                  string    `json:"role,omitempty"`
	CreatedAt             time.Time `json:"created_at"`
	PasswordLastSetAt     time.Time `json:"password_last_set_at"`
	LastLogin             time.Time `json:"last_login"`
	ForcePasswordChange   bool      `json:"force_password_change,omitempty"`
	FailedLoginCount      int       `json:"failed_login_count,omitempty"`
	Locked                bool      `json:"locked,omitempty"`
	LockedUntil           time.Time `json:"locked_until"`
	MFASecret             string    `json:"-"`
	MFASetupComplete      bool      `json:"mfa_setup_complete,omitempty"`
	RecentHashedPasswords []string  `json:"recent_hashed_passwords,omitempty"`
	MFAImage              []byte    `json:"-"`
	Timezone              string    `json:"timezone"`
}
