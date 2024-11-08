package initializers

import (
	"log"

	"github.com/jackc/pgx/v5"

	"food-diary/internal/db"
)

func ConnectToDB(logger *log.Logger) *pgx.Conn {
	logger.Println("Initialize DB Connection")
	db := db.Connect()
	return db
}
