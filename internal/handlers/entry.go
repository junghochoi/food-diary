package handlers

import (
	"fmt"
	"log"
	"net/http"

	"food-diary/internal/helpers"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, r *http.Request) {
	mockData := map[string]string{
		"id":       "entry1",
		"mealType": "Breakfast",
		"item":     "Apple",
		"desc":     "very healthy apple",
	}

	err := helpers.WriteJson(w, http.StatusOK, mockData, http.Header{})
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
	err = h.conn.QueryRow(req.Context(), query, input.Title, input.Foods, input.FoodsDescription, input.Rating, input.RatingDescription).
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
