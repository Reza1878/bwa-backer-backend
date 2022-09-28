package transaction

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

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

func (r *repositoryImpl) FindByTransactionCode(code string) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("code = ?", code).Find(&transaction).Error

	return transaction, err
}

func (r *repositoryImpl) Save(transaction Transaction) (Transaction, error) {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	code, err := r.generateTransactionCode(tx)

	transaction.Code = code
	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	err = r.db.Create(&transaction).Error

	if err != nil {
		tx.Rollback()
		return transaction, err
	}

	tx.Commit()
	return transaction, nil
}

func (r *repositoryImpl) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error

	return transaction, err
}

func (r *repositoryImpl) generateTransactionCode(tx *gorm.DB) (string, error) {
	var sequence TransactionNumberSequence

	err := tx.FirstOrInit(&sequence, TransactionNumberSequence{
		Year: time.Now().Year(),
	}).Error

	if err != nil {
		return "", err
	}
	sequence.Seq = sequence.Seq + 1

	code := fmt.Sprintf("TRX-BCKR-%d%04d", sequence.Year, sequence.Seq)

	err = tx.Save(&sequence).Error

	if err != nil {
		return "", err
	}

	return code, nil

}

func (r *repositoryImpl) FindTransactionSummary(dateStart string, dateEnd string, userId int) ([]TransactionSummary, error) {
	var transactionSummary []TransactionSummary
	query := r.db.Table("transactions").Select("DATE_FORMAT(created_at, '%b %Y') as period", "SUM(amount) as amount", "status", "user_id").Where("user_id = ?", userId).Where("status", "success")

	if dateStart != "" {
		fmt.Println(dateStart)
		query.Where("created_at >= ?", dateStart)
	}
	if dateEnd != "" {
		query.Where("created_at <= ?", dateEnd)
	}
	query.Group("period, status, user_id")
	err := query.Find(&transactionSummary).Error
	if err != nil {
		return transactionSummary, err
	}
	return transactionSummary, nil
}
