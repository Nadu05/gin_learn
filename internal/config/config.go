package config

import "os"

type Config struct {
	PORT string
	DB   string
}

func LoadConfig() Config {
	return Config{
		PORT: os.Getenv("PORT"),
		DB:   os.Getenv("DATABASE_URL"),
	}
}
