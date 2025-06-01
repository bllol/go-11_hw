package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName string
	Version string
	Port    string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found")
	}

	AppConfig = Config{
		AppName: GetEnv("APP_NAME", "GoLang"),
		Version: GetEnv("VERSION", "1.0.0"),
		Port:    GetEnv("PORT", "7070"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
