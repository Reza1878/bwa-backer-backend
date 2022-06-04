package transaction

import (
	"bwa-backer/campaign"
	"time"
)

type CampaignTransactionResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionResponse {
	response := CampaignTransactionResponse{}
	response.Id = transaction.ID
	response.Name = transaction.User.Name
	response.Amount = transaction.Amount
	response.CreatedAt = transaction.CreatedAt

	return response
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionResponse {
	responses := []CampaignTransactionResponse{}

	for _, t := range transactions {
		responses = append(responses, FormatCampaignTransaction(t))
	}

	return responses
}

type UserTransactionResponse struct {
	ID        int                             `json:"id"`
	Amount    int                             `json:"amount"`
	Status    string                          `json:"status"`
	CreatedAt time.Time                       `json:"created_at"`
	Campaign  UserTransactionCampaignResponse `json:"campaign"`
}

type UserTransactionCampaignResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionResponse {
	response := UserTransactionResponse{}

	response.ID = transaction.ID
	response.Amount = transaction.Amount
	response.Status = transaction.Status
	response.CreatedAt = transaction.CreatedAt

	campaignFormat := campaign.FormatCampaign(transaction.Campaign)

	response.Campaign.Name = campaignFormat.Name
	response.Campaign.ImageUrl = campaignFormat.ImageUrl

	return response
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionResponse {
	responses := []UserTransactionResponse{}

	for _, transaction := range transactions {
		responses = append(responses, FormatUserTransaction(transaction))
	}

	return responses
}
