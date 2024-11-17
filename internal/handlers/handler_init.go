package handlers

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"food-diary/internal/config"
	"food-diary/internal/repository/interfaces"
)

type Handlers struct {
	conf *config.Config
	pool *pgxpool.Pool

	entryRepo repository.EntryRepository
}

func NewHandlers(
	cfg *config.Config,
	conn *pgxpool.Pool,
	entryRepo repository.EntryRepository,
) *Handlers {
	return &Handlers{conf: cfg, pool: conn, entryRepo: entryRepo}
}
