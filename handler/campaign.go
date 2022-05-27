package handler

import (
	"bwa-backer/campaign"
)

type CampaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *CampaignHandler {
	return &CampaignHandler{
		campaignService: campaignService,
	}
}
