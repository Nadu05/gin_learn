package services

import (
	"context"
	"encoding/json"
	"ginLearn/internal/models"
	"ginLearn/internal/repository"
	"time"

	"github.com/redis/go-redis/v9"
)

type UserService struct {
	repo  *repository.UserRepository
	cache *redis.Client
}

func NewUserService(repo *repository.UserRepository, cache *redis.Client) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache,
	}
}

func (s *UserService) GetUserByID(id string) (models.User, error) {

	ctx := context.Background()
	key := "user:" + id

	// 1. CHECK CACHE
	val, err := s.cache.Get(ctx, key).Result()
	if err == nil {
		var user models.User
		err := json.Unmarshal([]byte(val), &user)
		if err != nil {
			return models.User{}, err
		}
		return user, nil
	}

	// 2. DB CALL
	user, err := s.repo.GetByID(id)
	if err != nil {
		return user, err
	}

	// 3. STORE CACHE
	data, _ := json.Marshal(user)
	s.cache.Set(ctx, key, data, 5*time.Minute)

	return user, nil
}
