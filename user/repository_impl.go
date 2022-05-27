package user

import "gorm.io/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (repository *repositoryImpl) Save(user User) (User, error) {
	err := repository.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositoryImpl) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositoryImpl) FindById(ID int) (User, error) {
	var user User
	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositoryImpl) Update(user User) (User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
