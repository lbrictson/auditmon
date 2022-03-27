package storage

import (
	"context"
	"errors"
	"time"

	"github.com/lbrictson/auditmon/ent/user"

	"github.com/google/uuid"
	"github.com/lbrictson/auditmon/ent"
	"github.com/lbrictson/auditmon/pkg/models"
)

type UserStore struct {
	client *ent.Client
}

type NewUserStoreInput struct {
	EntClient *ent.Client
}

func MustNewUserStore(config NewUserStoreInput) *UserStore {
	s := UserStore{}
	s.client = config.EntClient
	return &s
}

func convertEntUserToModelUser(u ent.User) models.User {
	user := models.User{
		ID:                    u.ID.String(),
		Username:              u.Username,
		HashedPassword:        u.HashedPassword,
		Role:                  u.Role,
		CreatedAt:             u.CreatedAt,
		PasswordLastSetAt:     u.PasswordLastSetAt,
		LastLogin:             u.LastLogin,
		ForcePasswordChange:   u.ForcePasswordChange,
		FailedLoginCount:      u.FailedLogins,
		Locked:                u.Locked,
		LockedUntil:           u.LockedUntil,
		MFASecret:             u.MfaSecret,
		MFASetupComplete:      u.MfaSetupCompleted,
		RecentHashedPasswords: u.RecentPasswords,
		MFAImage:              u.MfaImage,
		Timezone:              u.Timezone,
	}
	return user
}

type CreateUserInput struct {
	Username       string
	HashedPassword string
	Role           string
	MFASecret      string
	MFAImage       []byte
}

func (s *UserStore) Create(ctx context.Context, input CreateUserInput) (*models.User, error) {
	u, err := s.client.User.Create().
		SetCreatedAt(time.Now()).
		SetID(uuid.New()).
		SetFailedLogins(0).
		SetForcePasswordChange(true).
		SetPasswordLastSetAt(time.Now()).
		SetLastLogin(time.Now()).
		SetHashedPassword(input.HashedPassword).
		SetLocked(false).
		SetLockedUntil(time.Now()).
		SetMfaSetupCompleted(false).
		SetUsername(input.Username).
		SetMfaSecret(input.MFASecret).
		SetRecentPasswords([]string{}).
		SetMfaImage(input.MFAImage).
		SetTimezone("UTC").
		SetRole(input.Role).Save(ctx)
	if err != nil {
		return nil, err
	}
	user := convertEntUserToModelUser(*u)
	return &user, nil
}

func (s *UserStore) All(ctx context.Context) ([]models.User, error) {
	allEntUsers, err := s.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var users []models.User
	for _, x := range allEntUsers {
		users = append(users, convertEntUserToModelUser(*x))
	}
	return users, nil
}

func (s *UserStore) GetByUsername(ctx context.Context, username string) (models.User, error) {
	u := models.User{}
	entUser, err := s.client.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return u, errors.New("not found")
	}
	return convertEntUserToModelUser(*entUser), nil
}

func (s *UserStore) Update(ctx context.Context, u models.User) (models.User, error) {
	entUser, err := s.client.User.UpdateOneID(uuid.MustParse(u.ID)).
		SetFailedLogins(u.FailedLoginCount).
		SetForcePasswordChange(u.ForcePasswordChange).
		SetPasswordLastSetAt(u.PasswordLastSetAt).
		SetLastLogin(u.LastLogin).
		SetHashedPassword(u.HashedPassword).
		SetLocked(u.Locked).
		SetLockedUntil(u.LockedUntil).
		SetMfaSetupCompleted(u.MFASetupComplete).
		SetMfaSecret(u.MFASecret).
		SetRecentPasswords(u.RecentHashedPasswords).
		SetMfaImage(u.MFAImage).
		SetTimezone(u.Timezone).
		Save(ctx)
	if err != nil {
		return models.User{}, err
	}
	return convertEntUserToModelUser(*entUser), nil
}
