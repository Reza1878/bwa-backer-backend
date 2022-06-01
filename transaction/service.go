package transaction

type Service interface {
	GetTransactionsByCampaignId(request GetCampaignTransactionRequest) ([]Transaction, error)
}
