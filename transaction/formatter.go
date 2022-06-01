package transaction

import "time"

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
