package auth

import "github.com/golang-jwt/jwt"

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
