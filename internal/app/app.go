package app

import (
	"ginLearn/internal/config"
	handlers "ginLearn/internal/handler"
	"ginLearn/internal/repository"
	"ginLearn/internal/services"

	"gorm.io/gorm"
)

type App struct {
	DB          *gorm.DB
	UserHandler *handlers.UserHandler
}

func NewApp() *App {

	db := config.NewDB()

	userRepo := repository.NewUserRepository(db)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	return &App{
		DB:          db,
		UserHandler: userHandler,
	}
}
