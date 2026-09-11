package applications

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"sync"
	"time"
)

var ErrInvalidApplication = errors.New("name and contact are required")

type Application struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"createdAt"`
}

type Store interface {
	Create(context.Context, Application) error
	Close()
}

type Notifier interface {
	Notify(context.Context, Application) error
}

type Service struct {
	store    Store
	notifier Notifier
	logger   *slog.Logger
}

func NewService(store Store, notifier Notifier, logger *slog.Logger) *Service {
	return &Service{store: store, notifier: notifier, logger: logger}
}

func (s *Service) Create(ctx context.Context, name, contact string) (Application, error) {
	application, err := New(name, contact, time.Now())
	if err != nil {
		return Application{}, err
	}
	if err := s.store.Create(ctx, application); err != nil {
		return Application{}, fmt.Errorf("save application: %w", err)
	}
	if err := s.notifier.Notify(ctx, application); err != nil {
		s.logger.Error("notify about application", "application_id", application.ID, "error", err)
	}
	return application, nil
}

type MemoryStore struct {
	mu    sync.Mutex
	items []Application
}

func (s *MemoryStore) Create(_ context.Context, application Application) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items = append(s.items, application)
	return nil
}

func (s *MemoryStore) Close() {}

type NoopNotifier struct{}

func (NoopNotifier) Notify(context.Context, Application) error { return nil }

func New(name, contact string, now time.Time) (Application, error) {
	name = strings.TrimSpace(name)
	contact = strings.TrimSpace(contact)
	if name == "" || contact == "" {
		return Application{}, ErrInvalidApplication
	}
	return Application{
		ID:        newID(now),
		Name:      name,
		Contact:   contact,
		CreatedAt: now.UTC(),
	}, nil
}

func newID(now time.Time) string {
	random := make([]byte, 6)
	if _, err := rand.Read(random); err != nil {
		return now.UTC().Format("20060102T150405.000000000")
	}
	return now.UTC().Format("20060102T150405") + "-" + hex.EncodeToString(random)
}
