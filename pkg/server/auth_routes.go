package server

import (
	"context"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/lbrictson/auditmon/pkg/models"

	"github.com/labstack/echo-contrib/session"

	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"

	"github.com/lbrictson/auditmon/pkg/auth"

	"github.com/labstack/echo/v4"
)

func (s *Server) viewLogout(c echo.Context) error {
	return s.destroySession(c, "/login")
}

func (s *Server) viewLoginPage(c echo.Context) error {
	showError := false
	showMFAError := false
	locked := false
	mfa_enabled := false
	if c.QueryParam("login_error") == "true" {
		showError = true
	}
	if c.QueryParam("mfa_error") == "true" {
		showMFAError = true
	}
	if c.QueryParam("locked") == "true" {
		locked = true
	}
	if c.QueryParam("mfa_enabled") == "true" {
		mfa_enabled = true
	}
	return c.Render(http.StatusOK, "login", map[string]any{
		"login_error": showError,
		"mfa_error":   showMFAError,
		"locked":      locked,
		"mfa_enabled": mfa_enabled,
	})
}

func (s *Server) formLogin(c echo.Context) error {
	type FormData struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}
	data := FormData{}
	c.Bind(&data)
	if data.Username == "" {
		return c.Redirect(http.StatusSeeOther, "/login?login_error=true")
	}
	if data.Password == "" {
		return c.Redirect(http.StatusSeeOther, "/login?login_error=true")
	}
	u, err := s.userStorage.GetByUsername(context.TODO(), data.Username)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login?login_error=true")
	}
	if !auth.ComparePassword(data.Password, u.HashedPassword) {
		return c.Redirect(http.StatusSeeOther, "/login?login_error=true")
	}
	if u.Locked {
		return c.Redirect(http.StatusSeeOther, "/login?locked=true")
	}
	s.createSession(SessionData{
		Username:              u.Username,
		Role:                  u.Role,
		UserID:                u.ID,
		SessionID:             uuid.New().String(),
		MFACompleted:          false,
		MFASetupRequired:      !u.MFASetupComplete,
		PasswordResetRequired: u.ForcePasswordChange,
	}, c)
	return c.Redirect(http.StatusSeeOther, "/")
}

func (s *Server) viewNewMFAPage(c echo.Context) error {
	showError := c.QueryParam("error")
	showErrorMessage := false
	if showError == "true" {
		showErrorMessage = true
	}
	sessionData, _, err := s.readSessionData(c)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	u, err := s.userStorage.GetByUsername(context.TODO(), sessionData.Username)
	if err != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.Render(http.StatusOK, "new_mfa", map[string]any{
		"MFASecret": u.MFASecret,
		"QRImage":   base64.StdEncoding.EncodeToString(u.MFAImage),
		"ShowError": showErrorMessage,
	})
}

func (s *Server) formNewMFA(c echo.Context) error {
	type FormData struct {
		Code string `form:"mfa_code"`
	}
	sessionData, _, err := s.readSessionData(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/new_mfa?error=true")
	}
	data := FormData{}
	c.Bind(&data)
	if data.Code == "" {
		log.Warnf("user %v entered no code for MFA setup", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/new_mfa?error=true")
	}
	u, err := s.userStorage.GetByUsername(context.TODO(), sessionData.Username)
	if err != nil {
		log.Warnf("user %v not found in database for MFA setup", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	if u.MFASetupComplete {
		log.Warnf("user %v tried to do MFA setup but MFA setup was already done", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	if auth.ValidateMFA(data.Code, u.MFASecret) == false {
		log.Warnf("user %v enetered incorrect MFA token for MFA setup", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/new_mfa?error=true")
	}
	u.MFASetupComplete = true
	s.userStorage.Update(context.TODO(), u)
	log.Debugf("%v completed MFA setup", sessionData.Username)
	sess, _ := session.Get("auditmon", c)
	sess.Options.MaxAge = -1
	sess.Options.Path = "/"
	sess.Save(c.Request(), c.Response())
	s.eventStorage.Create(context.TODO(), []models.Event{
		{
			EventName:       "successful_mfa_setup",
			EventTime:       time.Now(),
			Username:        u.Username,
			Resource:        "auditmon_user_" + u.Username,
			EventSource:     "auditmon_internal",
			SourceIPAddress: c.RealIP(),
			RequestID:       c.Request().Header.Get(echo.HeaderXRequestID),
			ReadOnly:        false,
			EventData: map[string]any{
				"role": u.Role,
				"mfa":  true,
			},
		},
	})
	return c.Redirect(http.StatusSeeOther, "/login?mfa_enabled=true")
}

func (s *Server) viewMFAPage(c echo.Context) error {
	_, _, err := s.readSessionData(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	return c.Render(http.StatusOK, "mfa", nil)
}

func (s *Server) formMFA(c echo.Context) error {
	type FormData struct {
		Code string `form:"mfa_code"`
	}
	sessionData, _, err := s.readSessionData(c)
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	data := FormData{}
	c.Bind(&data)
	if data.Code == "" {
		log.Warnf("user %v entered no code for MFA login", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/login?mfa_error=true")
	}
	u, err := s.userStorage.GetByUsername(context.TODO(), sessionData.Username)
	if err != nil {
		log.Warnf("user %v not found in database for MFA login", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/login")
	}
	if auth.ValidateMFA(data.Code, u.MFASecret) == false {
		log.Warnf("user %v enetered incorrect MFA token for MFA login", sessionData.Username)
		return c.Redirect(http.StatusSeeOther, "/login?mfa_error=true")
	}
	s.createSession(SessionData{
		Username:              u.Username,
		Role:                  u.Role,
		UserID:                u.ID,
		SessionID:             sessionData.SessionID,
		MFACompleted:          true,
		MFASetupRequired:      false,
		PasswordResetRequired: false,
		Timezone:              u.Timezone,
	}, c)
	s.eventStorage.Create(context.TODO(), []models.Event{
		{
			EventName:       "successful_login",
			EventTime:       time.Now(),
			Username:        u.Username,
			Resource:        "auditmon_user_" + u.Username,
			EventSource:     "auditmon_internal",
			SourceIPAddress: c.RealIP(),
			RequestID:       c.Request().Header.Get(echo.HeaderXRequestID),
			ReadOnly:        false,
			EventData: map[string]any{
				"role": u.Role,
				"mfa":  true,
			},
		},
	})
	return c.Redirect(http.StatusSeeOther, "/")
}
