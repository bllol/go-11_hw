package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port       string
	AdminName  string
	AdminEmail string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found, using default valuse")
	}

	AppConfig = Config{
		Port:       GetEnv("PORT", "7070"),
		AdminName:  GetEnv("ADMIN_NAME", "Admin"),
		AdminEmail: GetEnv("ADMIN_EMAIL", "Email"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
