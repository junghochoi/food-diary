package initializers

import (
	"log"

	"gorm.io/gorm"
	"food-diary/internal/db"
)

func ConnectToDB(logger *log.Logger) *gorm.DB{
	logger.Println("Initialize DB Connection")
	db := db.Connect()
  return db
}
