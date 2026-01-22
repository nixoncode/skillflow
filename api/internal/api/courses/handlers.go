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

	return response.Ok(c, "create a course", course)
}

func (ch *CourseHandler) ListCourses(c echo.Context) error {
	return response.Ok(c, "list of courses", nil)
}

func (ch *CourseHandler) GetCourse(c echo.Context) error {
	return response.Ok(c, "get a course", nil)
}

func (ch *CourseHandler) UpdateCourse(c echo.Context) error {
	return response.Ok(c, "update a course", nil)
}

func (ch *CourseHandler) DeleteCourse(c echo.Context) error {
	return response.Ok(c, "delete a course", nil)
}
