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
