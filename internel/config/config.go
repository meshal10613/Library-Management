package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Dsn  string

	JwtDuration  string
	JwtSecretKey string

	AdminName     string
	AdminEmail    string
	AdminPassword string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Port:         os.Getenv("PORT"),
		Dsn:          os.Getenv("DSN"),
		
		JwtDuration:  os.Getenv("JWT_DURATION"),
		JwtSecretKey: os.Getenv("JWT_SECRET_KEY"),

		AdminName:     os.Getenv("ADMIN_NAME"),
		AdminEmail:    os.Getenv("ADMIN_EMAIL"),
		AdminPassword: os.Getenv("ADMIN_PASSWORD"),
	}, nil
}
