package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
}

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db: db}
}

func (repository *RepositoryImpl) Save(user User) (User, error) {
	err := repository.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *RepositoryImpl) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
