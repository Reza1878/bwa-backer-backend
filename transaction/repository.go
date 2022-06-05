package transaction

import "gorm.io/gorm"

type Repository interface {
	FindByCampaignId(campaignId int) ([]Transaction, error)
	FindByUserId(userId int) ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	generateTransactionCode(tx *gorm.DB) (string, error)
}
