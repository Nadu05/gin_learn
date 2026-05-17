package services

import (
	"ginLearn/internal/models"
	"ginLearn/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	return s.repo.GetByID(id)
}
