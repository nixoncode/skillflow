package enrollments

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nixoncode/skillflow/core"
	"github.com/nixoncode/skillflow/internal/api/user"
	"github.com/nixoncode/skillflow/pkg/response"
)

type EnrollmentHandler struct {
	app   core.App
	store *EnrollmentStore
}

func (eh *EnrollmentHandler) Enroll(c echo.Context) error {

	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	claims := user.GetClaims(c)
	userID := claims.UserID

	if err := eh.store.enroll(userID, courseID); err != nil {
		if errors.Is(err, ErrAlreadyEnrolled) {
			return response.Error(c, http.StatusConflict, "already enrolled")
		}
		eh.app.Log().Error().Err(err).Msg("failed to enroll")
		return response.Error(c, http.StatusInternalServerError, "failed to enroll")
	}

	return response.Created(c, "enrolled", map[string]int64{"course_id": courseID})
}

func (eh *EnrollmentHandler) ListMyEnrollments(c echo.Context) error {
	claims := user.GetClaims(c)
	userID := claims.UserID

	out, err := eh.store.listByUser(userID)
	if err != nil {
		eh.app.Log().Error().Err(err).Msg("failed to list enrollments")
		return response.Error(c, http.StatusInternalServerError, "failed to list enrollments")
	}

	return response.Ok(c, "enrollments", out)
}

func (eh *EnrollmentHandler) Unenroll(c echo.Context) error {
	courseID, err := echo.PathParam[int64](c, "courseId")
	if err != nil {
		return response.BadRequest(c, "invalid course id")
	}

	claims := user.GetClaims(c)
	userID := claims.UserID

	if err := eh.store.delete(userID, courseID); err != nil {
		eh.app.Log().Error().Err(err).Msg("failed to unenroll")
		return response.Error(c, http.StatusInternalServerError, "failed to unenroll")
	}

	return response.Ok(c, "unenrolled", nil)
}
