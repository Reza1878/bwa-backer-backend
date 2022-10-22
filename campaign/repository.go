package campaign

type Repository interface {
	FindAll(request GetCampaignsRequest) ([]Campaign, error)
	FindByUserID(UserID int) ([]Campaign, error)
	FindByID(campaignID int) (Campaign, error)
	FindCampaignBySlug(slug string) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)
	Update(campaign Campaign) (Campaign, error)
	SaveImage(campaignImage CampaignImage) (CampaignImage, error)
	FindImageById(imageId int) (CampaignImage, error)
	FindImageByCampaign(campaignId int) ([]CampaignImage, error)
	DeleteImageById(imageId int) error
	UpdateImage(image CampaignImage) (CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}
