package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sportTrackerAPI/internal/user"
)

type Service struct {
	*user.Repository
}

func NewAuthService(repository *user.Repository) *Service {
	return &Service{repository}
}

func (service *Service) Register(email, password, userName string) (string, error) {
	existedUser, _ := service.Repository.FindByEmail(email)
	if existedUser != nil {
		return "", errors.New("ErrUserExisted")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		UserName:       userName,
		Email:          email,
		HashedPassword: string(hashedPassword)}

	_, err = service.Repository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}

func (service Service) Login(email, password string) (string, error) {
	existedUser, _ := service.Repository.FindByEmail(email)
	if existedUser == nil {
		return "", errors.New("ErrWrongCredentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.HashedPassword), []byte(password))
	if err != nil {
		return "", err
	}

	return existedUser.Email, nil
}
