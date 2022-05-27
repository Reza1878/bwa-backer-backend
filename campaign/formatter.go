package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	UserId           int    `json:"user_id"`
	Slug             string `json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.ID = campaign.ID
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.UserId = campaign.UserId
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter = append(campaignFormatter, FormatCampaign(campaign))
	}

	return campaignFormatter
}
