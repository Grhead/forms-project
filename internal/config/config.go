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
	redirectUrl  string
}
type DbConfig struct {
	fileName string
}

type ConfigProvider interface {
	LoadFormConfig() (FormConfig, error)
	LoadDbConfig() (DbConfig, error)
}
type DatabaseProvider interface {
	Connect(cfg DbConfig) (*gorm.DB, error)
}
type EnvConfigProvider struct{}
type DbSQLiteProvider struct{}

func (e *EnvConfigProvider) LoadFormConfig() (*FormConfig, error) {
	err := godotenv.Load("configs/.env.form")
	if err != nil {
		return nil, err
	}
	cfg := &FormConfig{
		clientID:     os.Getenv("CLIENT_ID"),
		clientSecret: os.Getenv("CLIENT_SECRET"),
		redirectUrl:  os.Getenv("REDIRECT_URL"),
	}
	return cfg, nil
}

func (e *EnvConfigProvider) LoadDBConfig() (*DbConfig, error) {
	err := godotenv.Load("configs/.env.database")
	if err != nil {
		return nil, err
	}
	cfg := &DbConfig{
		fileName: os.Getenv("FILE_NAME"),
	}
	return cfg, nil
}

func (s *DbSQLiteProvider) Connect(cfg *DbConfig) (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(cfg.fileName), &gorm.Config{})
}

func NewOAuth2Config(cfg *FormConfig) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.clientID,
		ClientSecret: cfg.clientSecret,
		Endpoint:     google.Endpoint,
		Scopes:       []string{forms.FormsBodyScope, forms.FormsResponsesReadonlyScope},
		RedirectURL:  cfg.redirectUrl,
	}
}
