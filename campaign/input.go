package campaign

import "bwa-backer/user"

type GetCampaignDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}

type GetCampaignsRequest struct {
	UserId int
	Limit  int
	Name   string
}

type CreateCampaignRequest struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CreateCampaignImageRequest struct {
	CampaignID int   `form:"campaign_id" binding:"required"`
	IsPrimary  *bool `form:"is_primary" binding:"required"`
	User       user.User
}
