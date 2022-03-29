package server

import (
	"fmt"
	"io/fs"
	"net/http"
	"time"

	"github.com/lbrictson/auditmon/web"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lbrictson/auditmon/pkg/storage"
)

type Server struct {
	eventStorage      *storage.EventStore
	userStorage       *storage.UserStore
	callBackURL       string
	port              int
	sessionSecret     string
	maxSessionSeconds int
	minPasswordLength int
	maxOldPasswordUse int
	maxFailedLogins   int
	lockoutSeconds    int
	maxQueryResults   int
}

type NewServerInput struct {
	EventStorage      *storage.EventStore
	UserStorage       *storage.UserStore
	Port              int
	CallBackURL       string
	SessionSecret     string
	MaxSessionSeconds int
	MinPasswordLength int
	MaxOldPasswordUse int
	MaxFailedLogins   int
	LockoutSeconds    int
	MaxQueryResults   int
}

func MustNewServer(config NewServerInput) *Server {
	s := Server{}
	s.eventStorage = config.EventStorage
	s.userStorage = config.UserStorage
	s.port = config.Port
	s.callBackURL = config.CallBackURL
	s.lockoutSeconds = config.LockoutSeconds
	s.maxFailedLogins = config.MaxFailedLogins
	s.maxSessionSeconds = config.MaxSessionSeconds
	s.sessionSecret = config.SessionSecret
	s.minPasswordLength = config.MinPasswordLength
	s.maxOldPasswordUse = config.MaxOldPasswordUse
	s.maxQueryResults = config.MaxQueryResults
	return &s
}

// RunServer starts a blocking server process
func (s *Server) RunServer() {
	e := echo.New()
	e.Use(middleware.RequestID())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(s.sessionSecret))))
	e.HideBanner = true
	e.Server.IdleTimeout = 30 * time.Second
	e.Server.ReadTimeout = 15 * time.Second
	e.Server.ReadHeaderTimeout = 10 * time.Second
	e.Renderer = MustNewRenderer()
	fSys, err := fs.Sub(web.Assets, "static")
	if err != nil {
		panic(err)
	}
	// Serve static files from virtual file system 'static' directory
	assetHandler := http.FileServer(http.FS(fSys))
	e.GET("/static/*", echo.WrapHandler(http.StripPrefix("/static/", assetHandler)))
	e.GET("/api/v1/healthz", s.heartbeatRoute)
	// Auth routes
	e.GET("/login", s.viewLoginPage)
	e.GET("/new_mfa", s.viewNewMFAPage)
	e.POST("/form/login", s.formLogin)
	e.POST("/form/new_mfa", s.formNewMFA)
	e.GET("/logout", s.viewLogout)
	e.GET("/change_password", s.viewLoginPage)
	e.GET("/mfa", s.viewMFAPage)
	e.POST("/form/mfa", s.formMFA)
	// App routes
	frontendAuthRequiredPages := e.Group("", s.frontendAuthenticationRequiredMiddleware)
	frontendAuthRequiredPages.GET("/", s.viewIndexPage)
	frontendAuthRequiredPages.POST("/", s.formIndexPage)
	frontendAuthRequiredPages.GET("/hook/event/:id", s.hookGetEventByID)
	// Profile routes
	frontendAuthRequiredPages.GET("/profile", s.viewProfilePage)
	frontendAuthRequiredPages.POST("/form/timezone", s.formTimezone)
	// HTML components
	components := e.Group("/component", s.frontendAuthenticationRequiredMiddleware)
	components.GET("/preferences.html", s.componentPreferences)
	components.GET("/filterby-datalist.html", s.componentRenderAutofill)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", s.port)))
}
