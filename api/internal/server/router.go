package server

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/internal/api/profile"
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

	// protected routes
	protectedRoutes := apiRoutes.Group("")
	protectedRoutes.Use(useAuth(s.app.Config().JWT.SecretKey))
	{
		profile.RegisterProfileRoutes(protectedRoutes, s.app)
	}
}
