package enrollments

import (
	"fmt"

	"github.com/nixoncode/skillflow/core"
)

type EnrollmentStore struct {
	app core.App
}

var ErrAlreadyEnrolled = fmt.Errorf("already enrolled")

func (es *EnrollmentStore) enroll(userID, courseID int64) error {
	// prevent duplicate
	var exists int
	if err := es.app.DB().Get(&exists, "SELECT COUNT(1) FROM enrollments WHERE user_id = ? AND course_id = ?", userID, courseID); err != nil {
		return err
	}
	if exists > 0 {
		return ErrAlreadyEnrolled
	}

	_, err := es.app.DB().Exec("INSERT INTO enrollments (user_id, course_id, enrolled_at) VALUES (?, ?, CURRENT_TIMESTAMP)", userID, courseID)
	return err
}

func (es *EnrollmentStore) listByUser(userID int64) ([]EnrollmentView, error) {
	var out []EnrollmentView
	query := `SELECT e.course_id, c.title as course_title, e.enrolled_at FROM enrollments e LEFT JOIN courses c ON c.id = e.course_id WHERE e.user_id = ? ORDER BY e.enrolled_at DESC`
	if err := es.app.DB().Select(&out, query, userID); err != nil {
		return nil, err
	}
	return out, nil
}

func (es *EnrollmentStore) delete(userID, courseID int64) error {
	_, err := es.app.DB().Exec("DELETE FROM enrollments WHERE user_id = ? AND course_id = ?", userID, courseID)
	return err
}
