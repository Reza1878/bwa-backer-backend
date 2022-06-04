package transaction

type Repository interface {
	FindByCampaignId(campaignId int) ([]Transaction, error)
	FindByUserId(userId int) ([]Transaction, error)
}
