package lessons

import (
	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
)

func RegisterLessonRoutes(e *echo.Group, app core.App) {
	handler := &LessonHandler{app: app}

	lessons := e.Group("/courses/:courseId/lessons")
	{
		lessons.GET("", handler.ListLessons)
		lessons.POST("", handler.CreateLesson)
		lessons.GET("/:lessonId", handler.GetLesson)
		lessons.PUT("/:lessonId", handler.UpdateLesson)
		lessons.DELETE("/:lessonId", handler.DeleteLesson)
	}
}
