package transaction

import "bwa-backer/user"

type GetCampaignTransactionRequest struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}
