package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
}
