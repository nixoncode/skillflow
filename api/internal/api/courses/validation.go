package courses

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
)

func (req *createCourseRequest) Validate() error {
	return v.ValidateStruct(req,
		v.Field(&req.Title, v.Required, v.Length(5, 200)),
		v.Field(&req.Description, v.Required, v.Length(10, 1000)),
		v.Field(&req.Thumbnail, v.Required),
		v.Field(&req.Price, v.When(req.Price != 0, v.Min(0.00))),
	)
}
