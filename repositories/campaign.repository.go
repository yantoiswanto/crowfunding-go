package repositories

import (
	"crowfunding/models"

	"gorm.io/gorm"
)

type CampaignRepository interface {
	FindAll() ([]models.Campaign, error)
	FindByUserID(userID int) ([]models.Campaign, error)
	FindByID(ID int) (models.Campaign, error)
	Save(campaign models.Campaign) (models.Campaign, error)
	Update(campaign models.Campaign) (models.Campaign, error)
	CreateImage(campaignImage models.CampaignImage) (models.CampaignImage, error)
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)
}

type campaignRepository struct {
	db *gorm.DB
}

func NewCampaignRepository(db *gorm.DB) *campaignRepository {
	return &campaignRepository{db}
}

func (r *campaignRepository) FindAll() ([]models.Campaign, error) {
	var campaigns []models.Campaign
	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (r *campaignRepository) FindByUserID(UserID int) ([]models.Campaign, error) {
	var campaigns []models.Campaign

	err := r.db.Where("user_id = ?", UserID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}

func (r *campaignRepository) FindByID(ID int) (models.Campaign, error) {
	var campaign models.Campaign
	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) Save(campaign models.Campaign) (models.Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (r *campaignRepository) Update(campaign models.Campaign) (models.Campaign, error) {
	err := r.db.Save(&campaign).Error
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (r *campaignRepository) CreateImage(campaignImage models.CampaignImage) (models.CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}

	return campaignImage, nil
}

func (r *campaignRepository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	err := r.db.Model(&models.CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
