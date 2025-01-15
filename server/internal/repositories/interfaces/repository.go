package repository

import (
	"food-diary/internal/models"
)

type EntryRepository interface {
	Create(entry *models.Entry) error
	Get(id string) (*models.Entry, error)
}
