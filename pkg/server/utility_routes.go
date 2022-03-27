package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) heartbeatRoute(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
