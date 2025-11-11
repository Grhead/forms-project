package config

import (
	"context"
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

func CreateAuthLink(conf *oauth2.Config) string {
	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func ExchangeToken(ctx context.Context, conf *oauth2.Config, code string, filename string) error {
	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return err
	}
	err = saveToken(token, filename)
	if err != nil {
		return err
	}
	return nil
}

func saveToken(token *oauth2.Token, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err = json.NewEncoder(f).Encode(token); err != nil {
		return err
	}
	return nil
}

func ReadToken(filename string) (*oauth2.Token, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	token := &oauth2.Token{}
	if err = json.NewDecoder(f).Decode(token); err != nil {
		return nil, err
	}
	return token, nil
}
