package service

import "food-diary/internal/models"

type EntryService interface {
	Create(input *models.Entry) (*models.Entry, error)
	GetEntryById(id string) (*models.Entry, error)
	// ListEntries([]models.Entry, error)
}
