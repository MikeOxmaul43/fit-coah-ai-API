package jwt

import "github.com/golang-jwt/jwt/v5"

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
		return j.Secret, nil
	})
	if err != nil {
		return false, nil
	}
	return token.Valid, claims
}
