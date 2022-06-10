package transaction

import "bwa-backer/user"

type GetCampaignTransactionRequest struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionRequest struct {
	Amount     int `json:"amount" binding:"required"`
	CampaignId int `json:"campaign_id" binding:"required"`
	User       user.User
}

type TransactionNotificationRequest struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
