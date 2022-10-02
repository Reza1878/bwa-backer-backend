package auth

import "gorm.io/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindByToken(token string) (Authentication, error) {
	var authentication Authentication
	err := r.db.Where("token = ?", token).Find(&authentication).Error

	return authentication, err
}

func (r *repositoryImpl) Save(authentication Authentication) (Authentication, error) {
	err := r.db.Save(&authentication).Error

	return authentication, err
}

func (r *repositoryImpl) DeleteByToken(token string) error {
	err := r.db.Where("token = ?", token).Delete(&Authentication{}).Error
	return err
}
