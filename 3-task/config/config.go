package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	AppName     string
	Version     string
	UserName    string
	GreetingMsg string
	AdminName   string
	AdminEmail  string
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("error loading .env file")
	}

	AppConfig = Config{
		Port:        GetEnv("PORT", "7070"),
		AppName:     GetEnv("APP_NAME", "AppName"),
		Version:     GetEnv("VERSION", "1.0.0"),
		UserName:    GetEnv("USER_NAME", "UserName"),
		GreetingMsg: GetEnv("GREETING_MSG", "Hello, World!"),
		AdminName:   GetEnv("ADMIN_NAME", "AdminName"),
		AdminEmail:  GetEnv("ADMIN_EMAIL", "AdminEmail"),
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
