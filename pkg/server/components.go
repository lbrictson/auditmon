package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ComponentPreferences struct {
	Timezone string
	Username string
	Role     string
}

func (s *Server) componentPreferences(c echo.Context) error {
	sessData, ok := c.Get("userData").(SessionData)
	if !ok {
		return c.Render(http.StatusOK, "error", map[string]any{
			"ErrorMessage": "Unable to determine preferences",
		})
	}
	data := ComponentPreferences{
		Timezone: sessData.Timezone,
		Username: sessData.Username,
		Role:     sessData.Role,
	}
	return c.Render(http.StatusOK, "preferences", map[string]any{
		"Data": data,
	})
}
