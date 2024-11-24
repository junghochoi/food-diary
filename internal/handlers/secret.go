package handlers

import (
	"food-diary/internal/response"
	"net/http"
)

func (h *Handlers) Secret(w http.ResponseWriter, r *http.Request) {
	if err := response.Success(w, "You got the secret", nil); err != nil {
		response.Error(w, http.StatusInternalServerError, "Error writing JSON", err)
	}
}
