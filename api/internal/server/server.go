package server

import (
	"github.com/labstack/echo/v4"
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

	s.setupRoutes()

	return s
}

func (s *Server) Start(addr string) error {
	s.app.Log().Info().Msgf("Starting server on %s", addr)
	return s.echo.Start(addr)
}
