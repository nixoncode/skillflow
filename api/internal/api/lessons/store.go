package lessons

func (lh *LessonHandler) createLesson(lesson *Lesson) error {
	query := `INSERT INTO lessons (course_id, title, description, content_path, created_at) VALUES (:course_id, :title, :description, :content_path, :created_at)`
	result, err := lh.app.DB().NamedExec(query, lesson)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	lesson.ID = id
	lh.app.Log().Info().Msgf("New lesson created with ID: %d", id)
	return nil
}

func (lh *LessonHandler) listLessonsByCourse(courseID int64) ([]Lesson, error) {
	var lessons []Lesson
	query := `SELECT id, course_id, title, description, content_path, created_at FROM lessons WHERE course_id = ? ORDER BY created_at`
	err := lh.app.DB().Select(&lessons, query, courseID)
	if err != nil {
		return nil, err
	}
	return lessons, nil
}

func (lh *LessonHandler) getLessonByID(courseID, lessonID int64) (*Lesson, error) {
	var lesson Lesson
	query := `SELECT id, course_id, title, description, content_path, created_at FROM lessons WHERE id = ? AND course_id = ?`
	err := lh.app.DB().Get(&lesson, query, lessonID, courseID)
	if err != nil {
		return nil, err
	}
	return &lesson, nil
}

func (lh *LessonHandler) updateLesson(lesson *Lesson) error {
	query := `UPDATE lessons SET title = :title, description = :description, content_path = :content_path WHERE id = :id AND course_id = :course_id`
	_, err := lh.app.DB().NamedExec(query, lesson)
	return err
}

func (lh *LessonHandler) deleteLesson(courseID, lessonID int64) error {
	query := `DELETE FROM lessons WHERE id = ? AND course_id = ?`
	result, err := lh.app.DB().Exec(query, lessonID, courseID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	lh.app.Log().Info().Msgf("Deleted lesson id=%d course=%d rows=%d", lessonID, courseID, rows)
	return nil
}
