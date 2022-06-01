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
	User       user.User
	Campaign   campaign.Campaign
}
