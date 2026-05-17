package config

import (
	"ginLearn/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDB now returns a *gorm.DB pointer instead of a pgx pool
func NewDB() *gorm.DB {
	// 1. Get the connection string from your environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=admin password=admin1234 dbname=ginapi port=5432 sslmode=disable"
	}

	// 2. Open the GORM connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	log.Println("DB connection success")

	// Register all your models here so GORM builds them on startup
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Database auto-migration failed: %v", err)
	}
	log.Println("Database migration completed successfully! Tables are ready.")

	return db
}
