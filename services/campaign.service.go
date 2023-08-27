package services

import (
	"crowfunding/models"
	"crowfunding/repositories"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]models.Campaign, error)
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
