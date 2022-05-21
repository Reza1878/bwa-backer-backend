package user

import "gorm.io/gorm"

type RepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
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

func (r *RepositoryImpl) FindById(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *RepositoryImpl) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
