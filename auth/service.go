package auth

import "github.com/golang-jwt/jwt"

type Service interface {
	GenerateToken(userID int) (string, error)
	GenerateRefreshToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	ValidateRefreshToken(token string) (*jwt.Token, error)
	GetRefreshToken(token string) (Authentication, error)
	DeleteRefreshToken(token string) error
}
