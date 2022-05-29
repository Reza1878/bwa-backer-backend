package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

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

func (service *serviceImpl) GetCampaign(input GetCampaignDetailRequest) (Campaign, error) {
	campaign, err := service.repository.FindByID(input.ID)

	return campaign, err
}

func (s *serviceImpl) CreateCampaign(request CreateCampaignRequest) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = request.Name
	campaign.ShortDescription = request.ShortDescription
	campaign.Description = request.Description
	campaign.GoalAmount = request.GoalAmount
	campaign.Perks = request.Perks
	slugCandidate := fmt.Sprintf("%s %d", request.Name, request.User.Id)
	campaign.Slug = slug.Make(slugCandidate)

	existingCampaign, _ := s.repository.FindCampaignBySlug(campaign.Slug)
	if existingCampaign.ID != 0 {
		return campaign, errors.New("campaign name must be unique")
	}

	campaign.UserId = request.User.Id

	newCampaign, err := s.repository.Save(campaign)

	return newCampaign, err
}
