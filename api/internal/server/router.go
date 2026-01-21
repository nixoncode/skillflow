package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

func (s *Server) setupRoutes() {
	s.echo.GET("/", func(c echo.Context) error {
		return response.Ok(c, "Welcome to SkillFlow API", nil)
	})

	apiRoutes := s.echo.Group("/api")
	{
		user.RegisterUserRoutes(apiRoutes, s.app)
	}
}
