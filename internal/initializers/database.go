package initializers

import (
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"food-diary/internal/db"
)

func ConnectToDB(logger *log.Logger) *pgxpool.Pool {
	logger.Println("Initialize DB Connection")
	db := db.Connect()
	return db
}
