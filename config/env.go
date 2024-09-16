package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Env struct {
	ENV  string
	PORT string
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}
	return defaultValue
}

func loadEnv() (*Env, error) {
	logger := GetLogger("Env")
	logger.Info("loading .env file...")
	err := godotenv.Load()

	if err != nil {
		logger.Errorf("error loading .env file: %s", err.Error())
		return nil, errors.New("error loading .env file")
	}
	envVars := Env{
		ENV:  getEnv("ENV", "development"),
		PORT: getEnv("PORT", "8080"),
	}

	envValidator := validator.New()
	err = envValidator.Struct(envVars)
	if err != nil {
		logger.Errorf("error validating .env file: %s", err.Error())
		return nil, err
	}
	logger.Info(".env file loaded")
	return &envVars, nil
}
