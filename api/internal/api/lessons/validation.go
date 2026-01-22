package lessons

import v "github.com/go-ozzo/ozzo-validation/v4"

func (r *createLessonRequest) Validate() error {
	return v.ValidateStruct(r,
		v.Field(&r.CourseID, v.Required),
		v.Field(&r.Title, v.Required, v.Length(3, 200)),
		v.Field(&r.Description, v.Required, v.Length(10, 2000)),
		v.Field(&r.ContentPath, v.Required, v.Length(1, 255)),
	)
}
