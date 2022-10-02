package auth

import (
	"bwa-backer/helper"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

var SECRET_KEY = []byte(helper.GetDotEnvVariable("JWT_SECRET_KEY"))
var SECRET_REFRESH_TOKEN_KEY = []byte(helper.GetDotEnvVariable("JWT_REFRESH_TOKEN_KEY"))

func (s *serviceImpl) GenerateRefreshToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["data"] = map[string]interface{}{
		"user_id": userID,
	}
	claim["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_REFRESH_TOKEN_KEY)

	if err != nil {
		return "", err
	}

	_, err = s.repository.Save(Authentication{Token: signedToken})

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (s *serviceImpl) DeleteRefreshToken(token string) error {
	authentication, err := s.repository.FindByToken(token)
	if err != nil {
		return err
	}
	if authentication.ID == 0 {
		return errors.New("refresh token not found")
	}
	err = s.repository.DeleteByToken(token)
	return err
}

func (s *serviceImpl) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["data"] = map[string]interface{}{
		"user_id": userID,
	}
	claim["iat"] = time.Now().Unix()
	claim["exp"] = time.Now().Add(time.Second * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (c *serviceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
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
