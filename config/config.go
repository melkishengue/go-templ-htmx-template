package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host             string
	User             string
	Database         string
	Password         string
	Port             string
	AdditionalParams string
}

type Config struct {
	Environment string
	Port        string
	Database    DatabaseConfig
}

var config *Config

func GetConfig() Config {
	if config == nil {
		config = NewConfig()
	}

	return *config
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Environment: withDefault(os.Getenv("ENVIRONMENT"), "development"),
		Port:        withDefault(os.Getenv("PORT"), "3000"),
		Database: DatabaseConfig{
			Host:             required("DB_HOST"),
			User:             required("DB_USER"),
			Password:         required("DB_PASSWORD"),
			Database:         required("DB_NAME"),
			Port:             withDefault("DB_PORT", "5432"),
			AdditionalParams: os.Getenv("DB_ADDITIONAL_PARAMS"),
		},
	}
}

func required(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatalf("Could not load variable %s. Aborting...", key)
		os.Exit(1)
	}

	return val
}

func withDefault[T ~string](val T, defaultValue T) T {
	if val == "" {
		return defaultValue
	}

	return val
}
