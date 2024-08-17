package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/devfranklindiaz/notion-iol-integration/domain/service"
)

type NotionHandler struct {
	service *service.NotionService
}

func NewNotionHandler(service *service.NotionService) *NotionHandler {
	return &NotionHandler{service: service}
}

func (h *NotionHandler) Connect(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	if url == "" {
		http.Error(w, "url is required", http.StatusBadRequest)
		return
	}

	resp, err := h.service.Connect(context.Background(), url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": resp.Status,
	})
}
