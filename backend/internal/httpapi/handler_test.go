package httpapi_test

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"dereza-stories/backend/internal/applications"
	"dereza-stories/backend/internal/httpapi"
)

func TestCreateApplication(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	service := applications.NewService(&applications.MemoryStore{}, applications.NoopNotifier{}, logger)
	handler := httpapi.New(service, logger)
	request := httptest.NewRequest(http.MethodPost, "/api/applications", bytes.NewBufferString(`{"name":"Иван","contact":"@ivan"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusAccepted {
		t.Fatalf("expected %d, got %d: %s", http.StatusAccepted, response.Code, response.Body.String())
	}
}

func TestCreateApplicationRejectsEmptyFields(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	service := applications.NewService(&applications.MemoryStore{}, applications.NoopNotifier{}, logger)
	handler := httpapi.New(service, logger)
	request := httptest.NewRequest(http.MethodPost, "/api/applications", bytes.NewBufferString(`{"name":"","contact":""}`))
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected %d, got %d", http.StatusUnprocessableEntity, response.Code)
	}
}
