package transaction

import (
	"bwa-backer/campaign"
	"bwa-backer/user"
	"time"
)

type Transaction struct {
	ID         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PaymentUrl string
	User       user.User
	Campaign   campaign.Campaign
}

type TransactionSummary struct {
	Period string `json:"period"`
	Amount int    `json:"amount"`
}

type TransactionNumberSequence struct {
	ID   int
	Year int
	Seq  int
}
