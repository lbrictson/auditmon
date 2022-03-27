package server

import (
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

// frontendAuthenticationRequiredMiddleware will validate a user is logged in and has completed any needed MFA and
// mandatory password reset steps
func (s *Server) frontendAuthenticationRequiredMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		redirectCode := http.StatusTemporaryRedirect
		if strings.ToLower(c.Request().Method) == "post" {
			redirectCode = http.StatusSeeOther
		}
		sessData, messages, err := s.readSessionData(c)
		if err != nil {
			log.Debug("redirecting to login page because user has no session")
			return c.Redirect(redirectCode, "/login")
		}
		if sessData.MFASetupRequired {
			log.Debugf("redirecting %v to new_mfa because mfa setup required", sessData.Username)
			return c.Redirect(redirectCode, "/new_mfa")
		}
		if !sessData.MFACompleted {
			log.Debugf("redirecting %v to mfa page because mfa has not been entered yet", sessData.Username)
			return c.Redirect(redirectCode, "/mfa")
		}
		c.Set("userData", sessData)
		c.Set("infoMessage", messages.InfoMessage)
		c.Set("errorMessage", messages.ErrorMessage)
		// set request id
		reqID := c.Request().Header.Get(echo.HeaderXRequestID)
		c.Set("requestID", reqID)
		log.Debugf("user %v passed auth steps in middleware, allowing to proceed", sessData.Username)
		return next(c)
	}
}
