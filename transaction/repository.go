package transaction

type Repository interface {
	FindByCampaignId(campaignId int) ([]Transaction, error)
}
