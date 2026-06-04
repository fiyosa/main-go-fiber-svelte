package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_Env         string
	APP_Locale      string
	APP_Secret      string
	APP_JwtDuration string
	APP_DbUrl       string
	APP_ApiUrl      string
	APP_Port        string
)

func setup() {
	APP_Env = getEnv("APP_ENV", "local")
	APP_Locale = getEnv("APP_LOCALE", "en")
	APP_Secret = getEnv("APP_SECRET", "secret")
	APP_JwtDuration = getEnv("APP_JWT_DURATION", "1d")
	APP_DbUrl = getEnv("DB_URL", "postgresql://postgres:password@localhost:5432/portfolio")
	APP_ApiUrl = getEnv("API_URL", "http://localhost:8000")
	APP_Port = getEnv("APP_PORT", "8000")
}

func InitConfigApp() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error loading .env file: %v \n\n", err.Error())
		os.Exit(1)
	}
	setup()
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
