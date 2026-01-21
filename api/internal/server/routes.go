package server

import "github.com/labstack/echo/v4"

func (s *Server) setupRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome to SkillFlow API")
	})
}
