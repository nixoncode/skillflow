package courses

func (ch *CourseHandler) createCourse(course *Course) error {
	query := "INSERT INTO courses (title,user_id, description, thumbnail, price, is_published, created_at) VALUES (:title, :user_id, :description, :thumbnail, :price, :is_published, :created_at)"

	result, err := ch.app.DB().NamedExec(query, course)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	ch.app.Log().Info().Msgf("New course created with ID: %d", id)

	course.ID = id
	return nil

}

func (ch *CourseHandler) listCourses() ([]Course, error) {
	var courses []Course
	query := "SELECT id, user_id, title, description, thumbnail, price, is_published, created_at FROM courses"
	err := ch.app.DB().Select(&courses, query)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (ch *CourseHandler) getCourseByID(id int64) (*Course, error) {
	var course Course
	query := "SELECT id, user_id, title, description, thumbnail, price, is_published, created_at FROM courses WHERE id = ?"
	err := ch.app.DB().Get(&course, query, id)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (ch *CourseHandler) updateCourse(course *Course) error {
	query := "UPDATE courses SET title = :title, description = :description, thumbnail = :thumbnail, price = :price, is_published = :is_published WHERE id = :id"
	_, err := ch.app.DB().NamedExec(query, course)
	return err
}

func (ch *CourseHandler) deleteCourse(id int64) error {
	query := "DELETE FROM courses WHERE id = ?"
	result, err := ch.app.DB().Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	ch.app.Log().Info().Msgf("Deleted course ID %d, rows affected: %d", id, rowsAffected)
	return nil
}
