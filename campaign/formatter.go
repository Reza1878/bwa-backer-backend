package campaign

import (
	"strings"
)

type CampaignResponse struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignResponse {
	response := CampaignResponse{}
	response.Name = campaign.Name
	response.ShortDescription = campaign.ShortDescription
	response.ID = campaign.ID
	response.CurrentAmount = campaign.CurrentAmount
	response.GoalAmount = campaign.GoalAmount
	response.UserId = campaign.UserId
	response.Slug = campaign.Slug

	for _, image := range campaign.CampaignImages {
		if image.IsPrimary {
			response.ImageUrl = image.FileName
		}
	}

	return response
}

func FormatCampaigns(campaigns []Campaign) []CampaignResponse {
	responses := []CampaignResponse{}

	for _, campaign := range campaigns {
		responses = append(responses, FormatCampaign(campaign))
	}

	return responses
}

type CampaignDetailResponse struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	ShortDescription string                `json:"short_description"`
	ImageUrl         string                `json:"image_url"`
	GoalAmount       int                   `json:"goal_amount"`
	CurrentAmount    int                   `json:"current_amount"`
	UserId           int                   `json:"user_id"`
	Description      string                `json:"description"`
	BackerCount      int                   `json:"backer_count"`
	Slug             string                `json:"slug"`
	Perks            []string              `json:"perks"`
	User             CampaignDetailUser    `json:"user"`
	Images           []CampaignDetailImage `json:"images"`
}

type CampaignDetailUser struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignDetailImage struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailResponse {
	response := CampaignDetailResponse{}
	response.CurrentAmount = campaign.CurrentAmount
	response.BackerCount = campaign.BackerCount
	response.GoalAmount = campaign.GoalAmount
	response.Description = campaign.Description
	response.Name = campaign.Name
	response.ShortDescription = campaign.ShortDescription
	response.Slug = campaign.Slug
	response.UserId = campaign.UserId
	response.Id = campaign.ID

	response.User = CampaignDetailUser{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.Avatar,
	}

	images := []CampaignDetailImage{}

	for _, image := range campaign.CampaignImages {
		images = append(images, CampaignDetailImage{
			ImageUrl:  image.FileName,
			IsPrimary: image.IsPrimary,
		})
	}

	response.Images = images

	for _, image := range campaign.CampaignImages {
		if image.IsPrimary {
			response.ImageUrl = image.FileName
			break
		}
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	response.Perks = perks

	return response
}
