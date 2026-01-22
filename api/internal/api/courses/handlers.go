package courses

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

func (ch *CourseHandler) CreateCourse(c echo.Context) error {

	claims := user.GetClaims(c)
	ch.app.Log().Info().Msgf("CreateCourse endpoint hit by user ID: %d", claims.UserID)
	if claims.Role != "instructor" {
		return response.Error(c, 403, "only instructors can create courses")
	}

	var req createCourseRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "invalid request payload")
	}

	if err := req.Validate(); err != nil {
		return response.ValidationError(c, err)
	}

	course := &Course{
		Title:       req.Title,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Price:       req.Price,
		UserID:      claims.UserID,
		IsPublished: req.IsPublished,
		CreatedAt:   time.Now(),
	}

	if err := ch.createCourse(course); err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to create course")
		return response.Error(c, 500, "failed to create course")
	}

	return response.Ok(c, "course created successfully", course)
}

func (ch *CourseHandler) ListCourses(c echo.Context) error {

	//TODO: handle instructor vs student view on published courses and pagination
	coursers, err := ch.listCourses()
	if err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to list courses")
		return response.Error(c, 500, "failed to list courses")
	}

	return response.Ok(c, "list of courses", coursers)
}

func (ch *CourseHandler) GetCourse(c echo.Context) error {

	id, err := echo.PathParam[int64](c, "id")
	if err != nil {
		ch.app.Log().Error().Err(err).Str("id", c.Param("id")).Msg("invalid course ID")
		return response.BadRequest(c, "invalid course ID")
	}

	course, err := ch.getCourseByID(id)
	if err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to get course by id")
		return response.Error(c, 500, "failed to get course")
	}

	return response.Ok(c, "get a course", course)
}

func (ch *CourseHandler) UpdateCourse(c echo.Context) error {

	claims := user.GetClaims(c)
	ch.app.Log().Info().Msgf("UpdateCourse endpoint hit by user ID: %d", claims.UserID)
	if claims.Role != "instructor" {
		return response.Error(c, 403, "only instructors can update courses")
	}

	id, err := echo.PathParam[int64](c, "id")
	if err != nil {
		ch.app.Log().Error().Err(err).Str("id", c.Param("id")).Msg("invalid course ID")
		return response.BadRequest(c, "invalid course ID")
	}

	var req createCourseRequest
	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "invalid request payload")
	}

	// get course to verify ownership
	courseData, err := ch.getCourseByID(id)
	if err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to get course by id")
		return response.Error(c, 500, "failed to get course")
	}

	if claims.UserID != courseData.UserID {
		return response.Error(c, 403, "you can only update your own courses")
	}

	if err := req.Validate(); err != nil {
		return response.ValidationError(c, err)
	}

	course := &Course{
		ID:          id,
		UserID:      claims.UserID,
		Title:       req.Title,
		Description: req.Description,
		Thumbnail:   req.Thumbnail,
		Price:       req.Price,
		IsPublished: req.IsPublished,
	}

	if err := ch.updateCourse(course); err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to update course")
		return response.Error(c, 500, "failed to update course")
	}

	return response.Ok(c, "course updated successfully", course)
}

func (ch *CourseHandler) DeleteCourse(c echo.Context) error {
	claims := user.GetClaims(c)
	ch.app.Log().Info().Msgf("DeleteCourse endpoint hit by user ID: %d", claims.UserID)
	if claims.Role != "instructor" {
		return response.Error(c, 403, "only instructors can delete courses")
	}

	id, err := echo.PathParam[int64](c, "id")
	if err != nil {
		ch.app.Log().Error().Err(err).Str("id", c.Param("id")).Msg("invalid course ID")
		return response.BadRequest(c, "invalid course ID")
	}

	// get course to verify ownership
	course, err := ch.getCourseByID(id)
	if err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to get course by id")
		return response.Error(c, 500, "failed to get course")
	}

	if claims.UserID != course.UserID {
		return response.Error(c, 403, "you can only delete your own courses")
	}

	if err := ch.deleteCourse(id); err != nil {
		ch.app.Log().Error().Err(err).Msg("failed to delete course")
		return response.Error(c, 500, "failed to delete course")
	}

	return response.Ok(c, "course deleted successfully", nil)
}
