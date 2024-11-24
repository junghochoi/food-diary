package routes

import (
	"github.com/go-chi/chi/v5"

	"food-diary/internal/config"
	"food-diary/internal/handlers"
	auth "food-diary/internal/middlewares"
)

func InitializeRoutes(conf *config.Config, h *handlers.Handlers) chi.Router {
	r := chi.NewRouter()
	r.Get("/v1/healthcheck", h.Healthcheck)

	// ENTRY Routes
	r.Get("/v1/entry", h.GetEntry)
	r.Post("/v1/entry", h.CreateEntry)

	r.With(auth.AuthMiddleware(conf.JWTSecret)).Get("/v1/secret", h.Secret)

	return r
}
