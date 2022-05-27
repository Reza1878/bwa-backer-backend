package campaign

type serviceImpl struct {
	repository Repository
}

func NewService(repository Repository) *serviceImpl {
	return &serviceImpl{
		repository: repository,
	}
}

func (service *serviceImpl) GetCampaignList() ([]Campaign, error) {
	campaign, err := service.repository.FindAll()

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}
