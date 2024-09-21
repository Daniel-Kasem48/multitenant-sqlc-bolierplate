package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBHost     string
	JWTSecret  string

	AuthChallengeSecret string
	DeviceSecret        string
}

func LoadConfig() *Config {
	config := Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),

		AuthChallengeSecret: os.Getenv("AUTH_CHALLENGE_SECRET"),
		DeviceSecret:        os.Getenv("DEVICE_SECRET"),
	}

	// Debug logging
	fmt.Printf("Loaded Config: %+v\n", config)

	return &config
}
