package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"sportTrackerAPI/internal/user"
	"sportTrackerAPI/pkg/jwt"
	"time"
)

type Service struct {
	*user.Repository
	RedisRepository *Repository
}

func NewAuthService(repository *user.Repository, redisRepository *Repository) *Service {
	return &Service{
		Repository:      repository,
		RedisRepository: redisRepository}
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

func (service *Service) Login(email, password, secret string) (*LoginResponse, error) {
	existedUser, _ := service.Repository.FindByEmail(email)
	if existedUser == nil {
		return nil, errors.New("ErrWrongCredentials")
	}

	err := bcrypt.CompareHashAndPassword([]byte(existedUser.HashedPassword), []byte(password))
	if err != nil {
		return nil, err
	}
	accessToken, refreshToken, accessExp, refreshExp, err := jwt.GenerateTokens(secret, email)
	if err != nil {
		return nil, err
	}
	err = service.RedisRepository.Set(email, refreshToken, time.Until(refreshExp))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  accessExp,
		RefreshExpires: refreshExp,
	}, nil
}

func (service *Service) Refresh(refreshToken, secret string) (*RefreshResponse, error) {
	isValid, claims := jwt.NewJWT(secret).Parse(refreshToken)
	if !isValid {
		return nil, errors.New("Invalid refresh token")
	}

	storedToken, err := service.RedisRepository.Get(claims.Email)
	if err != nil || storedToken != refreshToken {
		return nil, errors.New("token not found or expired")
	}

	accessToken, refreshToken, accessExp, refreshExp, err := jwt.GenerateTokens(secret, claims.Email)
	if err != nil {
		return nil, err
	}

	err = service.RedisRepository.Set(claims.Email, refreshToken, time.Until(refreshExp))
	if err != nil {
		return nil, err
	}
	return &RefreshResponse{
		AccessToken:    accessToken,
		RefreshToken:   refreshToken,
		AccessExpires:  accessExp,
		RefreshExpires: refreshExp,
	}, nil
}

func (service *Service) Logout(email string) error {
	return service.RedisRepository.Delete(email)
}
