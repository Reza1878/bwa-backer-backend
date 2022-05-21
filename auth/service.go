package auth

type Service interface {
	GenerateToken(userID int) (string, error)
}
