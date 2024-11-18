package service

import (
	"food-diary/internal/models"
	"food-diary/internal/repository/interfaces"
	"food-diary/internal/services/interfaces"
)

type entryService struct {
	repo repository.EntryRepository
}

func NewEntryService(repo repository.EntryRepository) service.EntryService {
	return &entryService{repo: repo}
}

func (s *entryService) Create(entry *models.Entry) (*models.Entry, error) {
	err := s.repo.Create(entry)
	if err != nil {
		return nil, err
	}
	return entry, nil
}

func (s *entryService) GetEntryById(id string) (*models.Entry, error) {
	entry, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
