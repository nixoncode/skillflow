package progress

import "time"

type Progress struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id" db:"user_id"`
	CourseID    int64     `json:"course_id" db:"course_id"`
	LessonID    int64     `json:"lesson_id" db:"lesson_id"`
	IsCompleted bool      `json:"is_completed" db:"is_completed"`
	CompletedAt time.Time `json:"completed_at" db:"completed_at"`
}

type CourseProgressView struct {
	CourseID           int64   `json:"course_id"`
	CompletedLessons   int64   `json:"completed_lessons"`
	TotalLessons       int64   `json:"total_lessons"`
	ProgressPercent    float64 `json:"progress_percent"`
	CompletedLessonIDs []int64 `json:"completed_lesson_ids"`
}
