package transaction

import "gorm.io/gorm"

type repositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repositoryImpl {
	return &repositoryImpl{db: db}
}

func (r *repositoryImpl) FindByCampaignId(campaignId int) ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Where("campaign_id = ?", campaignId).Preload("User").Order("id desc").Find(&transactions).Error

	return transactions, err
}

func (r *repositoryImpl) FindByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Where("user_id = ?", userId).Preload("Campaign.CampaignImages").Find(&transactions).Order("id desc").Error

	return transactions, err
}

func (r *repositoryImpl) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}
