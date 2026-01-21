package user

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (r *RegisterRequest) Validate() error {
	return v.ValidateStruct(r,
		v.Field(&r.Email, v.Required, v.Length(5, 100), is.Email),
		v.Field(&r.Password, v.Required, v.Length(8, 100)),
		v.Field(&r.Role, v.When(r.Role != "", v.In("student", "instructor").Error("role must be either a student or instructor"))),
	)
}

func (r *LoginRequest) Validate() error {
	return v.ValidateStruct(r,
		v.Field(&r.Email, v.Required, is.Email),
		v.Field(&r.Password, v.Required),
	)
}
