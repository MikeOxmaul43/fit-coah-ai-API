package auth

type RegisterRequest struct {
	UserName string `json:"user-name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}
