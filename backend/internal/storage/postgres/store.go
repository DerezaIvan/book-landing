package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"dereza-stories/backend/internal/applications"
)

type Store struct{ pool *pgxpool.Pool }

func Open(ctx context.Context, databaseURL string) (*Store, error) {
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create postgres pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping postgres: %w", err)
	}
	return &Store{pool: pool}, nil
}

func (s *Store) Create(ctx context.Context, application applications.Application) error {
	_, err := s.pool.Exec(ctx, `
		INSERT INTO applications (id, name, contact, created_at)
		VALUES ($1, $2, $3, $4)
	`, application.ID, application.Name, application.Contact, application.CreatedAt)
	if err != nil {
		return fmt.Errorf("insert application: %w", err)
	}
	return nil
}

func (s *Store) Close() { s.pool.Close() }
