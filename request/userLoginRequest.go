package request

type UserLoginRequest struct {
	Mail     string `validate:"required,email"`
	Password string `validate:"required"`
}
