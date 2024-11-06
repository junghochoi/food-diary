package handlers

import (
	"gorm.io/gorm"

	"food-diary/internal/config"
)

type Handlers struct {
	conf *config.Config
	db   *gorm.DB
}

func NewHandlers(cfg *config.Config, db *gorm.DB) *Handlers {
	return &Handlers{conf: cfg, db: db}
}
