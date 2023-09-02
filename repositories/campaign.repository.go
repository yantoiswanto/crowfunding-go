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
