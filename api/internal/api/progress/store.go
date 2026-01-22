package progress

import (
	"github.com/nixoncode/skillflow/core"
)

type ProgressStore struct {
	app core.App
}

func (ps *ProgressStore) markCompleted(userID, courseID, lessonID int64) error {
	var exists int
	if err := ps.app.DB().Get(&exists, "SELECT COUNT(1) FROM progress WHERE user_id = ? AND course_id = ? AND lesson_id = ?", userID, courseID, lessonID); err != nil {
		return err
	}
	if exists > 0 {
		_, err := ps.app.DB().Exec("UPDATE progress SET is_completed = TRUE, completed_at = CURRENT_TIMESTAMP WHERE user_id = ? AND course_id = ? AND lesson_id = ?", userID, courseID, lessonID)
		return err
	}

	_, err := ps.app.DB().Exec("INSERT INTO progress (user_id, course_id, lesson_id, is_completed, completed_at) VALUES (?, ?, ?, TRUE, CURRENT_TIMESTAMP)", userID, courseID, lessonID)
	return err
}

func (ps *ProgressStore) getCourseProgress(userID, courseID int64) (*CourseProgressView, error) {
	var total int64
	if err := ps.app.DB().Get(&total, "SELECT COUNT(1) FROM lessons WHERE course_id = ?", courseID); err != nil {
		return nil, err
	}

	var completed int64
	if err := ps.app.DB().Get(&completed, "SELECT COUNT(1) FROM progress WHERE user_id = ? AND course_id = ? AND is_completed = TRUE", userID, courseID); err != nil {
		return nil, err
	}

	var ids []int64
	if err := ps.app.DB().Select(&ids, "SELECT lesson_id FROM progress WHERE user_id = ? AND course_id = ? AND is_completed = TRUE", userID, courseID); err != nil {
		return nil, err
	}

	percent := 0.0
	if total > 0 {
		percent = (float64(completed) / float64(total)) * 100.0
	}

	return &CourseProgressView{
		CourseID:           courseID,
		CompletedLessons:   completed,
		TotalLessons:       total,
		ProgressPercent:    percent,
		CompletedLessonIDs: ids,
	}, nil
}
