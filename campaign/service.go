package campaign

type Service interface {
	GetCampaigns(input GetCampaignsRequest) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailRequest) (Campaign, error)
	CreateCampaign(request CreateCampaignRequest) (Campaign, error)
	UpdateCampaign(requestId GetCampaignDetailRequest, requestData CreateCampaignRequest) (Campaign, error)
	CreateCampaignImage(request CreateCampaignImageRequest, fileLocation string) (CampaignImage, error)
	DeleteCampaignImage(request DeleteCampaignImageRequest) (CampaignImage, error)
}
