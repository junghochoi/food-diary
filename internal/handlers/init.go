package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"food-diary/internal/config"
)

type Handlers struct {
	conf *config.Config
	conn *pgxpool.Pool
}

func NewHandlers(cfg *config.Config, conn *pgxpool.Pool) *Handlers {
	return &Handlers{conf: cfg, conn: conn}
}
