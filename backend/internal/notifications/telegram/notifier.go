package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"dereza-stories/backend/internal/applications"
)

type Notifier struct {
	endpoint string
	chatID   string
	client   *http.Client
}

func New(botToken, chatID string) *Notifier {
	return &Notifier{
		endpoint: "https://api.telegram.org/bot" + botToken + "/sendMessage",
		chatID:   chatID,
		client:   &http.Client{Timeout: 8 * time.Second},
	}
}

func (n *Notifier) Notify(ctx context.Context, application applications.Application) error {
	payload := struct {
		ChatID string `json:"chat_id"`
		Text   string `json:"text"`
	}{
		ChatID: n.chatID,
		Text:   fmt.Sprintf("Новая заявка Dereza Stories\nИмя: %s\nКонтакт: %s\nID: %s", application.Name, application.Contact, application.ID),
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("encode telegram message: %w", err)
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, n.endpoint, bytes.NewReader(body))
	if err != nil {
		return fmt.Errorf("create telegram request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := n.client.Do(request)
	if err != nil {
		return fmt.Errorf("send telegram message: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("telegram returned status %d", response.StatusCode)
	}
	return nil
}
