package services

import (
	"crowfunding/models"
	"crowfunding/repositories"
	"crowfunding/request"
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]models.Campaign, error)
	GetCampaignByID(input request.GetCampaignDeatilInput) (models.Campaign, error)
	CreateCampaign(input request.CreateCampaign) (models.Campaign, error)
}

type service struct {
	repository repositories.CampaignRepository
}

func NewCampaignService(repository repositories.CampaignRepository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]models.Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (s *service) GetCampaignByID(input request.GetCampaignDeatilInput) (models.Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampaign(input request.CreateCampaign) (models.Campaign, error) {
	campaign := models.Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}
