package routes

import (
	"github.com/go-chi/chi/v5"

	"food-diary/internal/handlers"
)

func InitializeRoutes(h *handlers.Handlers) chi.Router {
	r := chi.NewRouter()
	r.Get("/v1/healthcheck", h.Healthcheck)

	// ENTRY Routes
	r.Get("/v1/entry", h.GetEntry)
	r.Post("/v1/entry", h.CreateEntry)

	return r
}
