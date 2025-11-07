package config_test

import (
	"testing"
	"tusur-forms/internal/config"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func TestEnvConfigProvider_LoadFormConfig(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    *config.FormConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var e config.EnvConfigProvider
			got, gotErr := e.LoadFormConfig()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadFormConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadFormConfig() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("LoadFormConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvConfigProvider_LoadDBConfig(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    *config.DBConfig
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var e config.EnvConfigProvider
			got, gotErr := e.LoadDBConfig()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("LoadDBConfig() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("LoadDBConfig() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("LoadDBConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBSQLiteProvider_Connect(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cfg     *config.DBConfig
		want    *gorm.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var s config.DBSQLiteProvider
			got, gotErr := s.Connect(tt.cfg)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("Connect() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("Connect() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewOAuth2Config(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cfg  *config.FormConfig
		want *oauth2.Config
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := config.NewOAuth2Config(tt.cfg)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewOAuth2Config() = %v, want %v", got, tt.want)
			}
		})
	}
}
