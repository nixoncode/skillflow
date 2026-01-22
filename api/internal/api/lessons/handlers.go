package lessons

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

type LessonHandler struct {
	app core.App
}

func (lh *LessonHandler) CreateLesson(c echo.Context) error {

	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	var req createLessonRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "invalid request payload")
	}
	if err := req.Validate(); err != nil {
		return response.ValidationError(c, err)
	}

	claims := user.GetClaims(c)
	if claims.Role != "instructor" {
		return response.Error(c, http.StatusForbidden, "only instructors can create lessons")
	}

	lesson := &Lesson{
		CourseID:    courseID,
		Title:       req.Title,
		Description: req.Description,
		ContentPath: req.ContentPath,
		CreatedAt:   time.Now(),
	}

	if err := lh.createLesson(lesson); err != nil {
		lh.app.Log().Error().Err(err).Msg("failed to create lesson")
		return response.Error(c, http.StatusInternalServerError, "failed to create lesson")
	}

	return response.Created(c, "lesson created", lesson)
}

func (lh *LessonHandler) ListLessons(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	lessons, err := lh.listLessonsByCourse(courseID)
	if err != nil {
		lh.app.Log().Error().Err(err).Msg("failed to list lessons")
		return response.InternalServerError(c, "failed to list lessons")
	}

	return response.Ok(c, "lessons list", lessons)
}

func (lh *LessonHandler) GetLesson(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	lessonID, err := echo.PathParam[int64](c, "lessonId")
	if err != nil {
		return response.BadRequest(c, "invalid lesson id")
	}

	lesson, err := lh.getLessonByID(courseID, lessonID)
	if err != nil {
		lh.app.Log().Error().Err(err).Msg("failed to get lesson")
		return response.Error(c, http.StatusNotFound, "lesson not found")
	}

	return response.Ok(c, "lesson", lesson)
}

func (lh *LessonHandler) UpdateLesson(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}
	lessonID, err := echo.PathParam[int64](c, "lessonId")
	if err != nil {
		return response.BadRequest(c, "invalid lesson id")
	}

	claims := user.GetClaims(c)
	if claims.Role != "instructor" {
		return response.Error(c, http.StatusForbidden, "only instructors can update lessons")
	}

	var req createLessonRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "invalid request payload")
	}
	if err := req.Validate(); err != nil {
		return response.ValidationError(c, err)
	}

	lesson, err := lh.getLessonByID(courseID, lessonID)
	if err != nil {
		return response.NotFound(c, "lesson not found")
	}

	lesson.Title = req.Title
	lesson.Description = req.Description
	lesson.ContentPath = req.ContentPath

	if err := lh.updateLesson(lesson); err != nil {
		lh.app.Log().Error().Err(err).Msg("failed to update lesson")
		return response.InternalServerError(c, "failed to update lesson")
	}

	return response.Ok(c, "lesson updated", lesson)
}

func (lh *LessonHandler) DeleteLesson(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}
	lessonID, err := echo.PathParam[int64](c, "lessonId")
	if err != nil {
		return response.BadRequest(c, "invalid lesson id")
	}

	claims := user.GetClaims(c)
	if claims.Role != "instructor" {
		return response.Forbidden(c, "only instructors can delete lessons")
	}

	if err := lh.deleteLesson(courseID, lessonID); err != nil {
		lh.app.Log().Error().Err(err).Msg("failed to delete lesson")
		return response.InternalServerError(c, "failed to delete lesson")
	}

	return response.Ok(c, "lesson deleted", nil)
}
