# Dereza Stories

Локальный проект семейной студии персональных детских книг Ивана и Алины Дереза.

## Структура

- `frontend` - Svelte 5 + SvelteKit + TypeScript.
- `frontend/src/lib/components/sections` - самостоятельные секции главной страницы.
- `frontend/src/lib/actions` - общая логика появления элементов при скролле.
- `backend` - Go API заявок.
- `backend/internal/storage/postgres` - PostgreSQL-реализация `applications.Store`.
- `backend/internal/notifications/telegram` - уведомления о новой заявке.
- `backend/migrations` - SQL-миграции.

## Быстрый локальный запуск

По умолчанию backend работает с временным in-memory хранилищем и не требует секретов.

В первом терминале:

```sh
make dev-backend
```

Во втором терминале:

```sh
make dev-frontend
```

Главная страница откроется на `http://localhost:5173`. Vite проксирует `/api` в Go-сервис на `http://localhost:8080`.

## PostgreSQL

1. Скопируйте `backend/.env.example` в `backend/.env` и заполните `DATABASE_URL`.
2. Создайте таблицу:

```sh
DATABASE_URL='postgres://...' make migrate
```

3. Запустите API с постоянным хранилищем:

```sh
cd backend
set -a
source .env
set +a
STORAGE_DRIVER=postgres go run ./cmd/server
```

## Telegram

Заполните вместе `TELEGRAM_BOT_TOKEN` и `TELEGRAM_CHAT_ID`. Если обе переменные пустые, уведомления отключены. Ошибка Telegram не приводит к потере уже сохранённой заявки и фиксируется в журнале сервиса.

## Проверки

```sh
make check
make test
make build
```
