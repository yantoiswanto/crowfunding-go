package repository

import (
	"crowfunding/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user models.User) (models.User, error)
	FindByEmail(email string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(email string) (models.User, error) {
	var user models.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, nil
	}

	return user, nil
}
