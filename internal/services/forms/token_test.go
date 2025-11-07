package services_test

import (
	"context"
	"testing"
	"tusur-forms/internal/services/forms"

	"golang.org/x/oauth2"
)

func TestCreateAuthLink(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		conf *oauth2.Config
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := services.CreateAuthLink(tt.conf)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CreateAuthLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExchangeToken(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		conf     *oauth2.Config
		code     string
		filename string
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := services.ExchangeToken(context.Background(), tt.conf, tt.code, tt.filename)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ExchangeToken() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ExchangeToken() succeeded unexpectedly")
			}
		})
	}
}
