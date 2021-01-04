package routes

type UserSignupRequest struct {
	Name     string `validate:"required"`
	Mail     string `validate:"required,email"`
	Surname  string `validate:"required"`
	Password string `validate:"required"`
}
