package campaign

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (service *serviceImpl) GetCampaigns(userId int) ([]Campaign, error) {
	var campaigns []Campaign
	var err error

	if userId == 0 {
		campaigns, err = service.repository.FindAll()
	} else {
		campaigns, err = service.repository.FindByUserID(userId)
	}

	return campaigns, err
}

func (service *serviceImpl) GetCampaign(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := service.repository.FindByID(input.ID)

	return campaign, err
}
