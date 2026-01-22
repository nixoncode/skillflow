package progress

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
)

func RegisterProgressRoutes(e *echo.Group, app core.App) {
	handler := &ProgressHandler{app: app, store: &ProgressStore{app: app}}

	e.POST("/progress/courses/:courseId/lessons/:lessonId", handler.MarkLessonComplete)
	e.GET("/progress/courses/:courseId", handler.GetCourseProgress)
}
