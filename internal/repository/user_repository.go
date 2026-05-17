package repository

import (
	"ginLearn/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByID(id string) (models.User, error) {
	var user models.User

	err := r.db.First(&user, "id = ?", id).Error

	return user, err
}
