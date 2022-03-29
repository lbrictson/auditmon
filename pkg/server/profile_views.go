package server

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func (s *Server) viewProfilePage(c echo.Context) error {
	user, err := s.userStorage.GetByUsername(context.TODO(), c.Get("userData").(SessionData).Username)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.Render(http.StatusOK, "profile", map[string]any{
		"User": user,
	})
}

func (s *Server) formTimezone(c echo.Context) error {
	type FormData struct {
		Timezone string `form:"timezone"`
	}
	data := FormData{}
	c.Bind(&data)
	if data.Timezone == "" {
		return c.Redirect(http.StatusSeeOther, "/profile")
	}
	user, err := s.userStorage.GetByUsername(context.TODO(), c.Get("userData").(SessionData).Username)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	// Attempt to parse the zone first to make sure it is valid
	_, err = time.LoadLocation(data.Timezone)
	if err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	user.Timezone = data.Timezone
	s.userStorage.Update(context.TODO(), user)
	s.createSession(SessionData{
		Username:              user.Username,
		Role:                  user.Role,
		UserID:                user.ID,
		SessionID:             uuid.New().String(),
		MFACompleted:          user.MFASetupComplete,
		MFASetupRequired:      !user.MFASetupComplete,
		PasswordResetRequired: false,
		Timezone:              user.Timezone,
	}, c)
	return c.Redirect(http.StatusSeeOther, "/profile")
}
