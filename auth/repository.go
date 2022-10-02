package auth

type Repository interface {
	Save(authentication Authentication) (Authentication, error)
	FindByToken(token string) (Authentication, error)
	DeleteByToken(token string) error
}
