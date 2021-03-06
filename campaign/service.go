package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaign(input GetCampaignDetailRequest) (Campaign, error)
	CreateCampaign(request CreateCampaignRequest) (Campaign, error)
	UpdateCampaign(requestId GetCampaignDetailRequest, requestData CreateCampaignRequest) (Campaign, error)
	CreateCampaignImage(request CreateCampaignImageRequest, fileLocation string) (CampaignImage, error)
}
