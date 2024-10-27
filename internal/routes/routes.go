package routes

import (
	"github.com/go-chi/chi/v5"

  "food-diary/internal/handlers"
  "food-diary/internal/config"
)


func InitializeRoutes(conf *config.Config) chi.Router {
  r := chi.NewRouter()

  h := handlers.NewHandlers(conf)

  r.Get("/v1/healthcheck",h.Healthcheck)
  r.Get("/v1/entry", h.GetEntry)


  return r
}

