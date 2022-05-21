package auth

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type JwtService struct {
}

func NewService() *JwtService {
	return &JwtService{}
}

var SECRET_KEY = []byte("kucingsayabisabertelur")

func (s *JwtService) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (c *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
