package server

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type SessionData struct {
	Username              string `json:"username,omitempty"`
	Role                  string `json:"role,omitempty"`
	UserID                string `json:"user_id,omitempty"`
	SessionID             string `json:"session_id,omitempty"`
	MFACompleted          bool   `json:"mfa_completed,omitempty"`
	MFASetupRequired      bool   `json:"mfa_setup_required,omitempty"`
	PasswordResetRequired bool   `json:"password_reset_required,omitempty"`
}

type SessionMessages struct {
	ErrorMessage string `json:"error_message"`
	InfoMessage  string `json:"info_message"`
}

func (s *Server) createSession(input SessionData, c echo.Context) {
	sess, _ := session.Get("auditmon", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   s.maxSessionSeconds,
		HttpOnly: true,
	}
	data, _ := json.Marshal(&input)
	jsonString := string(data)
	sess.Values["data"] = jsonString
	sess.Save(c.Request(), c.Response())
	return
}

// readSession will return the data from the current session, an error will be returned if there is no data.  Reading
// the data from the current session will read and then remove any error and info messages
func (s *Server) readSessionData(c echo.Context) (SessionData, SessionMessages, error) {
	sess, _ := session.Get("auditmon", c)
	messages := SessionMessages{}
	sessionData := SessionData{}
	msg := sess.Values["infoMessage"]
	if msg != nil {
		messages.InfoMessage = msg.(string)
	}
	errMsg := sess.Values["errorMessage"]
	if errMsg != nil {
		messages.ErrorMessage = errMsg.(string)
	}
	data := sess.Values["data"]
	if data != nil {
		stringData := data.(string)
		json.Unmarshal([]byte(stringData), &sessionData)
	}
	sess.Values["infoMessage"] = nil
	sess.Values["errorMessage"] = nil
	sess.Options.Path = "/"
	sess.Options.MaxAge = s.maxSessionSeconds
	sess.Options.HttpOnly = true
	sess.Save(c.Request(), c.Response())
	if sessionData.Username == "" {
		return sessionData, messages, errors.New("no session")
	}
	return sessionData, messages, nil
}

func (s *Server) destroySession(c echo.Context, path string) error {
	sess, _ := session.Get("auditmon", c)
	sess.Options.MaxAge = -1
	sess.Options.Path = "/"
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusTemporaryRedirect, path)
}
