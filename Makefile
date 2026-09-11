.PHONY: dev-frontend dev-backend check test build migrate

dev-frontend:
	cd frontend && npm run dev

dev-backend:
	cd backend && go run ./cmd/server

check:
	cd frontend && npm run check

test:
	cd backend && go test ./...

migrate:
	cd backend && psql "$(DATABASE_URL)" -f migrations/001_create_applications.sql

build:
	cd frontend && npm run build
	mkdir -p backend/bin
	cd backend && go build -o bin/server ./cmd/server
