package request

import (
	"crowfunding/models"
)

type GetCampaignDeatilInput struct {
	ID int `uri:"id" binding:"required"`
}

type CreateCampaign struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             models.User
}

type CreateCampaignImage struct {
	CampaignID int  `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       models.User
}
