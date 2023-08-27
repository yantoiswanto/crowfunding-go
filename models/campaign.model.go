package models

import "time"

type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImage    []CampaignImage
}

type CampaignImage struct {
	ID         int
	CampaignID int
	FileName   string
	isPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
