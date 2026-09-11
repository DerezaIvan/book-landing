package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"dereza-stories/backend/internal/applications"
	"dereza-stories/backend/internal/config"
	"dereza-stories/backend/internal/httpapi"
	telegramnotification "dereza-stories/backend/internal/notifications/telegram"
	postgresstorage "dereza-stories/backend/internal/storage/postgres"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	settings, err := config.Load()
	if err != nil {
		logger.Error("load configuration", "error", err)
		os.Exit(1)
	}

	ctx := context.Background()
	store, err := buildStore(ctx, settings)
	if err != nil {
		logger.Error("initialize storage", "error", err)
		os.Exit(1)
	}
	defer store.Close()

	notifier := buildNotifier(settings)
	service := applications.NewService(store, notifier, logger)
	server := &http.Server{
		Addr:              settings.Address,
		Handler:           httpapi.New(service, logger),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	errorsChannel := make(chan error, 1)
	go func() {
		logger.Info("Dereza Stories API started", "address", settings.Address, "storage", settings.StorageDriver)
		errorsChannel <- server.ListenAndServe()
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
		shutdownContext, cancel := context.WithTimeout(context.Background(), settings.ShutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownContext); err != nil {
			logger.Error("graceful shutdown", "error", err)
		}
	case err := <-errorsChannel:
		if !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server stopped", "error", err)
			os.Exit(1)
		}
	}
}

func buildStore(ctx context.Context, settings config.Config) (applications.Store, error) {
	if settings.StorageDriver == "postgres" {
		return postgresstorage.Open(ctx, settings.DatabaseURL)
	}
	return &applications.MemoryStore{}, nil
}

func buildNotifier(settings config.Config) applications.Notifier {
	if settings.TelegramBotToken == "" {
		return applications.NoopNotifier{}
	}
	return telegramnotification.New(settings.TelegramBotToken, settings.TelegramChatID)
}
