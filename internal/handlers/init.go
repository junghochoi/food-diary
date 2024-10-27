package handlers

import (
  "food-diary/internal/config"
)

type Handlers struct {
  conf *config.Config
}

func NewHandlers(cfg *config.Config) *Handlers {
  return &Handlers{conf: cfg}
}
