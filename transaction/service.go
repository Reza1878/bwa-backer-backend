package transaction

type Service interface {
	GetTransactionsByCampaignId(request GetCampaignTransactionRequest) ([]Transaction, error)
	GetTransactionsByUserId(userId int) ([]Transaction, error)
	GetTransactionSummary(request GetTransactionSummaryRequest) ([]TransactionSummary, error)
	CreateTransaction(request CreateTransactionRequest) (Transaction, error)
	ProcessTransaction(request TransactionNotificationRequest) error
}
