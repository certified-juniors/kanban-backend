package handlers

import (
	"fmt"
	"net/http"
)

// Info Посмотреть версию сервиса
// @Summary Посмотреть версию сервиса
// @Description Посмотреть версию сервиса
// @Tags Help
// @Router /info [get]
// @Success 200 {string} string "OK"
func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, "Kanban 1.0")
}
