package handlers

import (
	"github.com/jackc/pgx/v5"

	"food-diary/internal/config"
)

type Handlers struct {
	conf *config.Config
	conn *pgx.Conn
}

func NewHandlers(cfg *config.Config, conn *pgx.Conn) *Handlers {
	return &Handlers{conf: cfg, conn: conn}
}
