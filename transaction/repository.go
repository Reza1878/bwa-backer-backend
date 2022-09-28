package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignId(campaignId int) ([]Transaction, error)
	FindByUserId(userId int) ([]Transaction, error)
	FindByTransactionCode(code string) (Transaction, error)
	FindTransactionSummary(dateStart string, dateEnd string, userId int) ([]TransactionSummary, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	generateTransactionCode(tx *gorm.DB) (string, error)
}
