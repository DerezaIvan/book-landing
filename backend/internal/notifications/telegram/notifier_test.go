package telegram

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"dereza-stories/backend/internal/applications"
)

func TestNotifySendsApplication(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			ChatID string `json:"chat_id"`
			Text   string `json:"text"`
		}
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Fatalf("decode: %v", err)
		}
		if payload.ChatID != "42" {
			t.Errorf("expected chat id 42, got %q", payload.ChatID)
		}
		if !strings.Contains(payload.Text, "Алина") || !strings.Contains(payload.Text, "@alina") {
			t.Errorf("unexpected text %q", payload.Text)
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	notifier := New("token", "42")
	notifier.endpoint = server.URL
	notifier.client = server.Client()
	err := notifier.Notify(context.Background(), applications.Application{ID: "app-1", Name: "Алина", Contact: "@alina"})
	if err != nil {
		t.Fatalf("notify: %v", err)
	}
}
