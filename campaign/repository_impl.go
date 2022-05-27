package campaign

import "gorm.io/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (repository *repositoryImpl) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := repository.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (repository *repositoryImpl) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := repository.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error

	return campaigns, err

}
