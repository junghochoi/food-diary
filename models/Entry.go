package models

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Entry struct {
	ID         string
	Title      string
	Foods      []string
	FoodDesc   string
	Rating     uint8
	RatingDesc string
}

// GetEntryByID retrieves an entry by its ID from the database
func GetEntryByID(ctx context.Context, pool *pgxpool.Pool, id string) (*Entry, error) {
	var entry Entry

	query := `
        SELECT id, title, foods, food_desc, rating, rating_desc
        FROM entries
        WHERE id = $1
    `

	err := pool.QueryRow(ctx, query, id).Scan(
		&entry.ID,
		&entry.Title,
		&entry.Foods,
		&entry.FoodDesc,
		&entry.Rating,
		&entry.RatingDesc,
	)
	if err != nil {
		return nil, err // return the error if no rows or any other error occurred
	}

	return &entry, nil
}

func CreateEntry(ctx context.Context, pool *pgxpool.Pool, entry *Entry) error {
	query := `
    INSERT INTO entries (title, foods, food_desc, rating, rating_desc)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
  `

	var entryId string

	err := pool.QueryRow(
		ctx,
		query,
		entry.Title,
		entry.Foods,
		entry.FoodDesc,
		entry.Rating,
		entry.RatingDesc,
	).Scan(&entryId)

	entry.ID = entryId

	return err
}
