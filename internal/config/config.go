package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	config := &Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("REDIRECT_URL"),
	}
	return config, nil
}
