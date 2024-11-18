package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"food-diary/internal/config"
	"food-diary/internal/services/interfaces"
)

type Handlers struct {
	conf *config.Config
	pool *pgxpool.Pool

	entryService service.EntryService
}

func NewHandlers(
	cfg *config.Config,
	conn *pgxpool.Pool,
	entryService service.EntryService,
) *Handlers {
	return &Handlers{conf: cfg, pool: conn, entryService: entryService}
}
