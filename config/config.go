package config

import (
	"log/slog"
	"os"
)

type AppConfig struct {
	Port        string
	PostgresDSN string
	Logger      *slog.Logger
}

func LoadConfig() *AppConfig {

	return &AppConfig{
		Port:        os.Getenv("PORT"),
		PostgresDSN: "host=localhost user=postgres password=pgpassword dbname=cubikdb port=5432 sslmode=disable",
	}
}
