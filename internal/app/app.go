package app

import (
	"context"
	"tusur-forms/internal/config"
)

func Run() {
	ctx := context.Background()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
}
