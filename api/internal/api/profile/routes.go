package profile

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

type ProfileHandler struct {
	app core.App
}

func RegisterProfileRoutes(e *echo.Group, app core.App) {
	handler := &ProfileHandler{app: app}

	profiles := e.Group("/profile")
	{
		profiles.GET("/me", handler.GetProfile)
	}
}

func (ph *ProfileHandler) GetProfile(c echo.Context) error {
	claims := user.GetClaims(c)
	if claims == nil {
		return response.Error(c, http.StatusUnauthorized, "Invalid token claims")
	}

	userID := claims.UserID

	return response.Ok(c, "get user profile", map[string]any{
		"user_id": userID,
		"role":    claims.Role,
	})
}
