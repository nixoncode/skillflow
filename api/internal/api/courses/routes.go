package courses

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
)

type CourseHandler struct {
	app core.App
}

func RegisterCourseRoutes(e *echo.Group, app core.App) {
	handler := &CourseHandler{app: app}

	courses := e.Group("/courses")
	{
		courses.GET("", handler.ListCourses)
		courses.POST("", handler.CreateCourse)
		courses.GET("/:id", handler.GetCourse)
		courses.PUT("/:id", handler.UpdateCourse)
		courses.DELETE("/:id", handler.DeleteCourse)
	}
}
