package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx/v5"

	"food-diary/internal/helpers"
	"food-diary/models"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, req *http.Request) {
	var entry models.Entry
	var input struct {
		Id string `json:"id"`
	}

	err := helpers.ReadJson(w, req, &input)

	query := `
    SELECT * 
    FROM entries
    WHERE ID = $1
  `
	err = h.pool.QueryRow(req.Context(), query, input.Id).Scan(
		&entry.ID,
		&entry.Title,
		&entry.Foods,
		&entry.FoodDesc, // pgx.NullString can handle NULL values
		&entry.Rating,
		&entry.RatingDesc,
	)
	if err != nil {
		if err == pgx.ErrNoRows {
			// No entry found for the provided ID
			log.Printf("No entry found with ID %v", input.Id)

			// Create and send a JSON response for no rows found
			response := map[string]interface{}{
				"success": false,
				"message": "No entry found",
			}
			helpers.WriteJson(w, http.StatusNotFound, response, http.Header{})
		} else {
			// General database query error
			log.Printf("Error scanning row: %v", err)

			// Create and send a JSON response for the database error
			response := map[string]interface{}{
				"success": false,
				"message": "Error retrieving entry",
			}
			helpers.WriteJson(w, http.StatusInternalServerError, response, http.Header{})
		}
		return
	}

	err = helpers.WriteJson(w, http.StatusOK, entry, http.Header{})
	if err != nil {
		http.Error(w, "Failed to Write Json Response", http.StatusInternalServerError)
	}
}

func (h *Handlers) CreateEntry(w http.ResponseWriter, req *http.Request) {
	var input struct {
		Title             string   `json:"title"`
		Foods             []string `json:"foods"`
		FoodsDescription  string   `json:"foodDesc"`
		Rating            uint8    `json:"rating"`
		RatingDescription string   `json:"ratingDesc"`
	}

	var output struct {
		Success      bool   `json:"success"`
		RowsAffected uint32 `json:"rowsAffected"`
		Entry        struct {
			Title             string   `json:"title"`
			Foods             []string `json:"foods"`
			FoodsDescription  string   `json:"foodDesc"`
			Rating            uint8    `json:"rating"`
			RatingDescription string   `json:"ratingDesc"`
		} `json:"entry"`
	}

	err := helpers.ReadJson(w, req, &input)
	if err != nil {
		log.Printf("CreateEntry Input error: %v", err)

		http.Error(w, fmt.Sprintf("Error reading input: %v", err), http.StatusBadRequest)
		return
	}

	query := `
    INSERT INTO entries (title, foods, food_desc, rating, rating_desc)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING id
  `

	var entryID string
	err = h.pool.QueryRow(req.Context(), query, input.Title, input.Foods, input.FoodsDescription, input.Rating, input.RatingDescription).
		Scan(&entryID)
	if err != nil {
		log.Printf("Error inserting entry: %v", err)
		output.Success = false
		helpers.WriteJson(w, http.StatusInternalServerError, output, http.Header{})
		return
	}

	output.Success = true
	output.Entry.Title = input.Title
	output.Entry.Foods = input.Foods
	output.Entry.FoodsDescription = input.FoodsDescription
	output.Entry.Rating = input.Rating
	output.Entry.RatingDescription = input.RatingDescription
	//
	err = helpers.WriteJson(w, http.StatusOK, output, http.Header{})
	if err != nil {
		http.Error(w, "Failed to Write Json Response", http.StatusInternalServerError)
	}
}
