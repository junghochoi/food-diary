package pgsql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"food-diary/internal/models"
	"food-diary/internal/repository/interfaces"
)

type entryPGRepository struct {
	db *pgxpool.Pool
}

func NewEntryRepository(db *pgxpool.Pool) repository.EntryRepository {
	return &entryPGRepository{db: db}
}

func (r *entryPGRepository) Create(entry *models.Entry) error {
	query := `
    INSERT INTO entries (title, foods, food_desc, rating, rating_desc)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
  `
	_, err := r.db.Exec(
		context.Background(),
		query,
		entry.Title,
		entry.Foods,
		entry.FoodDesc,
		entry.Rating,
		entry.RatingDesc,
	)

	return err
}

func (r *entryPGRepository) Get(id string) (*models.Entry, error) {
	var entry models.Entry
	err := r.db.QueryRow(context.Background(), "SELECT id, title, foods, food_desc, rating, rating_desc FROM entries WHERE id=$1", id).
		Scan(&entry.ID, &entry.Title, &entry.Foods, &entry.FoodDesc, &entry.Rating, &entry.RatingDesc)
	return &entry, err
}
