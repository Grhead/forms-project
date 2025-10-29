package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/forms/v1"
)

type config struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

type Provider interface {
	NewConfig() (*oauth2.Config, error)
}
type EnvProvider struct{}

func (e *EnvProvider) NewConfig() (*oauth2.Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	cfg := &config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectUrl:  os.Getenv("REDIRECT_URL"),
	}
	var config = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{forms.FormsBodyScope, forms.FormsResponsesReadonlyScope},
		RedirectURL:  cfg.RedirectUrl,
	}
	return config, nil
}
