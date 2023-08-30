package request

type GetCampaignDeatilInput struct {
	ID int `uri:"id" binding:"required"`
}
