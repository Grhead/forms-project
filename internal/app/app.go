package app

import (
	"context"
	"tusur-forms/internal/config"
)

func Run() error {
	ctx := context.Background()
	provider := &config.EnvProvider{}
	cfg, err := provider.NewFormConfig()
	if err != nil {
		return err
	}
	return nil
}
