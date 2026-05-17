package app

import (
	"ginLearn/internal/cache"
	"ginLearn/internal/config"
	handlers "ginLearn/internal/handler"
	"ginLearn/internal/repository"
	"ginLearn/internal/services"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type App struct {
	DB          *gorm.DB
	Redis       *redis.Client
	UserHandler *handlers.UserHandler
}

func NewApp() *App {

	rdb := cache.NewRedis()

	db := config.NewDB()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo, rdb)
	userHandler := handlers.NewUserHandler(userService)

	return &App{
		DB:          db,
		UserHandler: userHandler,
		Redis:       rdb,
	}
}
