package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
)

type UserHandler struct {
	app core.App
}

func RegisterUserRoutes(e *echo.Group, app core.App) {
	handler := &UserHandler{app: app}

	users := e.Group("/users")
	{
		users.GET("/me", handler.Profile)
	}

	auth := e.Group("/auth")
	{
		auth.POST("/login", handler.Login)
		auth.POST("/register", handler.Register)
		auth.POST("/logout", handler.Logout)
	}
}
