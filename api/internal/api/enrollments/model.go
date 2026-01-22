package enrollments

import "time"

type Enrollment struct {
	ID         int64     `json:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	CourseID   int64     `json:"course_id" db:"course_id"`
	EnrolledAt time.Time `json:"enrolled_at" db:"enrolled_at"`
}

type EnrollmentView struct {
	CourseID    int64     `json:"course_id"`
	CourseTitle *string   `json:"course_title,omitempty"`
	EnrolledAt  time.Time `json:"enrolled_at"`
}
