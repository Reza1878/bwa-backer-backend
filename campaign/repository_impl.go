package campaign

import (
	"errors"

	"gorm.io/gorm"
)

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (repository *repositoryImpl) FindAll(request GetCampaignsRequest) ([]Campaign, error) {
	var campaigns []Campaign

	query := repository.db.Preload("CampaignImages", "campaign_images.is_primary = true")

	if request.Limit != 0 {
		query.Limit(request.Limit)
	}

	if request.Name != "" {
		query.Where("name LIKE ?", "%"+request.Name+"%")
	}
	query.Order("id")

	err := query.Find(&campaigns).Error

	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (repository *repositoryImpl) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := repository.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = true").Find(&campaigns).Error

	return campaigns, err

}

func (repository *repositoryImpl) FindByID(campaignID int) (Campaign, error) {
	var campaign Campaign

	err := repository.
		db.
		Where("id = ?", campaignID).
		Preload("User").
		Preload("CampaignImages").
		Find(&campaign).Error

	return campaign, err
}

func (r *repositoryImpl) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *repositoryImpl) FindCampaignBySlug(slug string) (Campaign, error) {
	var campaign Campaign

	err := r.
		db.
		Where("slug = ?", slug).
		Preload("User").
		Preload("CampaignImages").
		Find(&campaign).Error

	return campaign, err
}

func (r *repositoryImpl) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error

	return campaign, err
}

func (r *repositoryImpl) SaveImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error

	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *repositoryImpl) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&CampaignImage{}).Where("campaign_id", campaignID).Update("is_primary", false).Error

	if err != nil {
		return false, err
	}
	return true, err
}

func (r *repositoryImpl) FindImageById(imageId int) (CampaignImage, error) {
	var image CampaignImage
	err := r.db.Model(&CampaignImage{}).Where("id = ?", imageId).First(&image).Error

	if image.ID == 0 {
		return image, errors.New("not found")
	}

	return image, err
}

func (r *repositoryImpl) DeleteImageById(imageId int) error {
	err := r.db.Where("id = ?", imageId).Preload("Campaign").Delete(&CampaignImage{}).Error

	return err
}

func (r *repositoryImpl) FindImageByCampaign(campaignId int) ([]CampaignImage, error) {
	var campaignImages []CampaignImage

	err := r.db.Where("campaign_id = ?", campaignId).Find(&campaignImages).Error

	return campaignImages, err
}

func (r *repositoryImpl) UpdateImage(image CampaignImage) (CampaignImage, error) {
	err := r.db.Save(&image).Error
	return image, err
}
