package config

import (
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/forms/v1"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type FormConfig struct {
	clientID     string
	clientSecret string
	redirectURL  string
}
type DBConfig struct {
	fileName string
}

type CfgProvider interface {
	LoadFormConfig() (FormConfig, error)
	LoadDbConfig() (DBConfig, error)
}
type DatabaseProvider interface {
	Connect(cfg DBConfig) (*gorm.DB, error)
}
type EnvConfigProvider struct{}
type DBSQLiteProvider struct{}

func (e *EnvConfigProvider) LoadFormConfig() (*FormConfig, error) {
	err := godotenv.Load("configs/.env.form")
	if err != nil {
		return nil, err
	}
	cfg := &FormConfig{
		clientID:     os.Getenv("CLIENT_ID"),
		clientSecret: os.Getenv("CLIENT_SECRET"),
		redirectURL:  os.Getenv("REDIRECT_URL"),
	}
	return cfg, nil
}

func (e *EnvConfigProvider) LoadDBConfig() (*DBConfig, error) {
	err := godotenv.Load("configs/.env.database")
	if err != nil {
		return nil, err
	}
	cfg := &DBConfig{
		fileName: os.Getenv("FILE_NAME"),
	}
	return cfg, nil
}

func (s *DBSQLiteProvider) Connect(cfg *DBConfig) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(cfg.fileName), &gorm.Config{})
}

func NewOAuth2Config(cfg *FormConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.clientID,
		ClientSecret: cfg.clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{forms.FormsBodyScope, forms.FormsResponsesReadonlyScope},
		RedirectURL:  cfg.redirectURL,
	}
}
