package lessons

import "time"

type Lesson struct {
	ID          int64     `json:"id"`
	CourseID    int64     `json:"course_id" db:"course_id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	ContentPath string    `json:"content_path" db:"content_path"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type createLessonRequest struct {
	CourseID    int64   `json:"course_id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	ContentPath string  `json:"content_path"`
}
