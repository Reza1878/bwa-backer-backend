package campaign

type Service interface {
	GetCampaignList() ([]Campaign, error)
}
