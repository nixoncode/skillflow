package progress

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

type ProgressHandler struct {
	app   core.App
	store *ProgressStore
}

func (ph *ProgressHandler) MarkLessonComplete(c echo.Context) error {

	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}
	lessonID, err := echo.PathParam[int64](c, "lessonId")
	if err != nil {
		return response.BadRequest(c, "invalid lesson id")
	}

	claims := user.GetClaims(c)
	userID := claims.UserID

	if err := ph.store.markCompleted(userID, courseID, lessonID); err != nil {
		ph.app.Log().Error().Err(err).Msg("failed to mark lesson complete")
		return response.Error(c, http.StatusInternalServerError, "failed to mark lesson complete")
	}

	return response.Created(c, "marked complete", map[string]int64{"lesson_id": lessonID, "course_id": courseID})
}

func (ph *ProgressHandler) GetCourseProgress(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	claims := user.GetClaims(c)
	userID := claims.UserID

	out, err := ph.store.getCourseProgress(userID, courseID)
	if err != nil {
		ph.app.Log().Error().Err(err).Msg("failed to fetch progress")
		return response.Error(c, http.StatusInternalServerError, "failed to fetch progress")
	}

	return response.Ok(c, "progress", out)
}
