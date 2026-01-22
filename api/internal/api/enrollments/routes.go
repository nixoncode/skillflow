package enrollments

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
)

func RegisterEnrollmentRoutes(e *echo.Group, app core.App) {
	handler := &EnrollmentHandler{app: app, store: &EnrollmentStore{app: app}}

	// course scoped enroll/unenroll
	e.POST("/courses/:courseId/enroll", handler.Enroll)
	e.DELETE("/courses/:courseId/enroll", handler.Unenroll)

	// user's enrollments
	e.GET("/users/me/enrollments", handler.ListMyEnrollments)
}
