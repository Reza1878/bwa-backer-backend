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

func (service *serviceImpl) GetCampaigns(input GetCampaignsRequest) ([]Campaign, error) {
	var campaigns []Campaign
	var err error

	if input.UserId == 0 {
		campaigns, err = service.repository.FindAll(input)
	} else {
		campaigns, err = service.repository.FindByUserID(input.UserId)
	}

	return campaigns, err
}

func (service *serviceImpl) GetCampaign(input GetCampaignDetailRequest) (Campaign, error) {
	campaign, err := service.repository.FindByID(input.ID)

	if campaign.ID == 0 {
		return campaign, errors.New("data not found")
	}

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

func (s *serviceImpl) UpdateCampaign(requestId GetCampaignDetailRequest, requestData CreateCampaignRequest) (Campaign, error) {
	campaign, err := s.repository.FindByID(requestId.ID)
	if err != nil {
		return campaign, err
	}
	campaign.Name = requestData.Name
	campaign.ShortDescription = requestData.ShortDescription
	campaign.Description = requestData.Description
	campaign.GoalAmount = requestData.GoalAmount
	campaign.Perks = requestData.Perks

	if campaign.UserId != requestData.User.Id {
		return campaign, errors.New("you are not allowed to update this campaign")
	}

	campaign.UserId = requestData.User.Id

	newCampaign, err := s.repository.Update(campaign)

	return newCampaign, err
}

func (s *serviceImpl) CreateCampaignImage(request CreateCampaignImageRequest, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindByID(request.CampaignID)

	campaignImage := CampaignImage{}
	campaignImage.CampaignID = request.CampaignID
	campaignImage.FileName = fileLocation
	campaignImage.IsPrimary = *request.IsPrimary

	if err != nil || campaign.ID == 0 {
		return campaignImage, errors.New("campaign not found")
	}

	if campaign.UserId != request.User.Id {
		return campaignImage, errors.New("you are not allowed to create image on this campaign")
	}

	if *request.IsPrimary {
		success, err := s.repository.MarkAllImagesAsNonPrimary(request.CampaignID)
		if err != nil || !success {
			return campaignImage, errors.New("internal server error")
		}
	}

	campaignImage, err = s.repository.SaveImage(campaignImage)

	return campaignImage, err
}

func (s *serviceImpl) DeleteCampaignImage(request DeleteCampaignImageRequest) (CampaignImage, error) {
	campaignImage, err := s.repository.FindImageById(request.CampaignImageId)
	if err != nil {
		return campaignImage, err
	}

	campaign, _ := s.repository.FindByID(campaignImage.CampaignID)

	if campaign.UserId != request.User.Id {
		return campaignImage, errors.New("you are not authorized to delete this resource")
	}

	campaignImages, err := s.repository.FindImageByCampaign(campaignImage.CampaignID)
	if err != nil {
		return campaignImage, err
	}

	if campaignImage.IsPrimary {
		for _, v := range campaignImages {
			if v.ID != campaignImage.ID {
				v.IsPrimary = true
				_, err = s.repository.UpdateImage(v)
				break
			}
		}
	}

	if err != nil {
		return campaignImage, err
	}

	err = s.repository.DeleteImageById(request.CampaignImageId)
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}
