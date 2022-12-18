package config

import (
	"github.com/gofiber/fiber/v2"
	"recipe.core/internal/storage"
)

type Config struct {
	*fiber.App
	Store *storage.Storage
}

func NewConfig() (*Config, error) {
	store, err := storage.New()
	if err != nil {
		return nil, err
	}

	return &Config{
		App:   fiber.New(),
		Store: store,
	}, nil
}
