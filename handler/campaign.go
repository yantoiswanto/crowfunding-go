package handler

import (
	"crowfunding/helper"
	"crowfunding/models"
	"crowfunding/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service services.CampaignService
}

func NewCampaignHandler(service services.CampaignService) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userID)
	if err != nil {
		response := helper.APIResponse("Error to get campaigns", http.StatusBadRequest, "error", models.FormatCampaigns(campaigns))

		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaigns", http.StatusOK, "success", models.FormatCampaigns(campaigns))

	c.JSON(http.StatusOK, response)

}
