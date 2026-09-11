package httpapi

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"dereza-stories/backend/internal/applications"
)

type Handler struct {
	service *applications.Service
	logger  *slog.Logger
}

func New(service *applications.Service, logger *slog.Logger) http.Handler {
	h := &Handler{service: service, logger: logger}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", h.health)
	mux.HandleFunc("POST /api/applications", h.createApplication)
	return withCommonHeaders(mux)
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) createApplication(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name    string `json:"name"`
		Contact string `json:"contact"`
	}
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, 1<<20))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&input); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}

	application, err := h.service.Create(r.Context(), input.Name, input.Contact)
	if errors.Is(err, applications.ErrInvalidApplication) {
		writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	if err != nil {
		h.logger.Error("save application", "error", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "could not save application"})
		return
	}

	h.logger.Info("application received", "id", application.ID, "name", application.Name)
	writeJSON(w, http.StatusAccepted, map[string]string{"id": application.ID, "status": "accepted"})
}

func withCommonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		slog.Error("encode response", "error", err)
	}
}
