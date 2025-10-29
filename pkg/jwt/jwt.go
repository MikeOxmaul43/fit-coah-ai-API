package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWT struct {
	Secret string
}

func NewJWT(secret string) *JWT {
	return &JWT{Secret: secret}
}

func (j *JWT) Create(data Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	s, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return s, nil
}

func (j *JWT) Parse(t string) (bool, *Claims) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(t, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return false, nil
	}
	return token.Valid, claims
}

func GenerateTokens(secret, email string) (accessToken string, refreshToken string, accessExp time.Time, refreshExp time.Time, err error) {
	accessExp = time.Now().Add(1 * time.Hour)
	claims := Claims{
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(accessExp), IssuedAt: jwt.NewNumericDate(time.Now())},
	}
	accessToken, err = NewJWT(secret).Create(claims)
	if err != nil {
		return
	}

	refreshExp = time.Now().Add(24 * 7 * time.Hour)
	claims = Claims{
		Email:            email,
		RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(refreshExp), IssuedAt: jwt.NewNumericDate(time.Now())},
	}
	refreshToken, err = NewJWT(secret).Create(claims)
	if err != nil {
		return
	}
	return
}
