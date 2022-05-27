package campaign

type Repository interface {
	FindAll() ([]Campaign, error)
	FindByUserID(UserID int) ([]Campaign, error)
}
