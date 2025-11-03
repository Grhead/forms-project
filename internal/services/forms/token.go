package internal

import (
	"context"
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

func createAuthLink(conf *oauth2.Config) string {
	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return authURL
}

func exchangeToken(ctx context.Context, conf *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func saveToken(token *oauth2.Token, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	f.Close()
	if err = json.NewEncoder(f).Encode(token); err != nil {
		return err
	}
	return nil
}

func readToken(filename string) (*oauth2.Token, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	f.Close()
	token := &oauth2.Token{}
	if err = json.NewDecoder(f).Decode(token); err != nil {
		return nil, err
	}
	return token, nil
}
