package services

import (
	"crowfunding/models"
	"crowfunding/repositories"
	"crowfunding/request"
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type CampaignService interface {
	GetCampaigns(userID int) ([]models.Campaign, error)
	GetCampaignByID(input request.GetCampaignDeatilInput) (models.Campaign, error)
	CreateCampaign(input request.CreateCampaign) (models.Campaign, error)
	UpdateCampaign(inputID request.GetCampaignDeatilInput, inputData request.CreateCampaign) (models.Campaign, error)
	SaveCampaignImage(input request.CreateCampaignImage, fileLocation string) (models.CampaignImage, error)
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

func (s *service) UpdateCampaign(inputID request.GetCampaignDeatilInput, inputData request.CreateCampaign) (models.Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("Not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil

}

func (s *service) SaveCampaignImage(input request.CreateCampaignImage, fileLocation string) (models.CampaignImage, error) {
	campaign, err := s.repository.FindByID(input.CampaignID)
	if err != nil {
		return models.CampaignImage{}, err
	}

	if campaign.UserID != input.User.ID {
		return models.CampaignImage{}, errors.New("Not an owner of the campaign")
	}

	isPrimary := 0
	if input.IsPrimary {
		isPrimary = 1

		_, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignID)
		if err != nil {
			return models.CampaignImage{}, err
		}
	}

	campaignImage := models.CampaignImage{}
	campaignImage.CampaignID = input.CampaignID
	campaignImage.IsPrimary = isPrimary
	campaignImage.FileName = fileLocation

	newCampaignImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newCampaignImage, err
	}

	return newCampaignImage, nil

}
