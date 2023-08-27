package services

import (
	"crowfunding/models"
	"crowfunding/repositories"
)

type CampaignService interface {
	FindCampaigns(userID int) ([]models.Campaign, error)
}

type campaignService struct {
	repository repositories.CampaignRepository
}

func NewServiceCampaign(campaignRepository repositories.CampaignRepository) *campaignService {
	return &campaignService{campaignRepository}
}

func (s *campaignService) FindCampaigns(userID int) ([]models.Campaign, error) {
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
