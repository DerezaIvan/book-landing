package config_test

import (
	"testing"

	"dereza-stories/backend/internal/config"
)

func TestLoadDefaultsToMemoryStorage(t *testing.T) {
	t.Setenv("PORT", "")
	t.Setenv("STORAGE_DRIVER", "")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	settings, err := config.Load()
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	if settings.StorageDriver != "memory" {
		t.Fatalf("expected memory storage, got %q", settings.StorageDriver)
	}
	if settings.Address != ":8080" {
		t.Fatalf("expected :8080, got %q", settings.Address)
	}
}

func TestPostgresRequiresDatabaseURL(t *testing.T) {
	t.Setenv("STORAGE_DRIVER", "postgres")
	t.Setenv("DATABASE_URL", "")
	t.Setenv("TELEGRAM_BOT_TOKEN", "")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	if _, err := config.Load(); err == nil {
		t.Fatal("expected configuration error")
	}
}

func TestTelegramSettingsMustBeConfiguredTogether(t *testing.T) {
	t.Setenv("STORAGE_DRIVER", "memory")
	t.Setenv("TELEGRAM_BOT_TOKEN", "token")
	t.Setenv("TELEGRAM_CHAT_ID", "")

	if _, err := config.Load(); err == nil {
		t.Fatal("expected configuration error")
	}
}
