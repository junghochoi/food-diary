package handlers

import (
  "food-diary/internal/config"
	"gorm.io/gorm"
)

type Handlers struct {
  conf *config.Config
  db *gorm.DB
}

func NewHandlers(cfg *config.Config, db *gorm.DB) *Handlers {
  return &Handlers{conf: cfg, db: db}
}
