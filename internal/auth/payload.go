package auth

import "time"

type RegisterRequest struct {
	UserName string `json:"user-name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginResponse struct {
	AccessToken    string    `json:"access-token"`
	RefreshToken   string    `json:"refresh-token"`
	AccessExpires  time.Time `json:"access-expires-in"`
	RefreshExpires time.Time `json:"refresh-expires-in"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh-token"`
}

type RefreshResponse struct {
	AccessToken    string    `json:"access-token"`
	RefreshToken   string    `json:"refresh-token"`
	AccessExpires  time.Time `json:"access-expires-in"`
	RefreshExpires time.Time `json:"refresh-expires-in"`
}
