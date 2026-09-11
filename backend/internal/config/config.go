package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Config struct {
	Address          string
	StorageDriver    string
	DatabaseURL      string
	TelegramBotToken string
	TelegramChatID   string
	ShutdownTimeout  time.Duration
}

func Load() (Config, error) {
	config := Config{
		Address:          ":" + valueOrDefault("PORT", "8080"),
		StorageDriver:    valueOrDefault("STORAGE_DRIVER", "memory"),
		DatabaseURL:      strings.TrimSpace(os.Getenv("DATABASE_URL")),
		TelegramBotToken: strings.TrimSpace(os.Getenv("TELEGRAM_BOT_TOKEN")),
		TelegramChatID:   strings.TrimSpace(os.Getenv("TELEGRAM_CHAT_ID")),
		ShutdownTimeout:  10 * time.Second,
	}

	if config.StorageDriver != "memory" && config.StorageDriver != "postgres" {
		return Config{}, fmt.Errorf("unsupported STORAGE_DRIVER %q", config.StorageDriver)
	}
	if config.StorageDriver == "postgres" && config.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required when STORAGE_DRIVER=postgres")
	}
	if (config.TelegramBotToken == "") != (config.TelegramChatID == "") {
		return Config{}, errors.New("TELEGRAM_BOT_TOKEN and TELEGRAM_CHAT_ID must be configured together")
	}
	return config, nil
}

func valueOrDefault(key, fallback string) string {
	if value := strings.TrimSpace(os.Getenv(key)); value != "" {
		return value
	}
	return fallback
}
