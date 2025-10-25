package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/oauth2"
)

func createAuthLink(conf *oauth2.Config) {
	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
}

func exchangeToken(ctx context.Context, conf *oauth2.Config, code string) (*oauth2.Token, error) {
	tok, err := conf.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok, nil
}

func saveToken(token *oauth2.Token) error {
	f, err := os.Create("token.json")
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to cache oauth token: %v", err)
		}
	}(f)
	err = json.NewEncoder(f).Encode(token)
	if err != nil {
		return err
	}
	return nil
}

func tokenFromFile(file string) *oauth2.Token {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Unable to open token file: %v", err)
		return nil
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Unable to cache token file: %v", err)
		}
	}(f)
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok
}
