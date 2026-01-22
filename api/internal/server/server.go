package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nixoncode/skillflow/core"
)

type Server struct {
	app  core.App
	echo *echo.Echo
}

func NewServer(app core.App) *Server {
	s := &Server{
		app:  app,
		echo: echo.New(),
	}

	s.echo.HideBanner = true
	s.echo.HidePort = true

	corsOrigin := app.Config().Server.CORSAllowedOrigins

	s.echo.Use(
		middleware.RequestID(),
		middleware.Recover(),
		middleware.Secure(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins:     []string{corsOrigin},
			AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
			AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
			AllowCredentials: true,
		}),
	)

	s.echo.Use(serverLogger(s))

	s.setupRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	s.app.Log().Info().Msgf("Starting server on %s", addr)
	return s.echo.Start(addr)
}
