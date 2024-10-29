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
		MealType    string   `json:"mealType"`
		Foods       []string `json:"item"`
		Description string   `json:"desc"`
	}

	err := helpers.ReadJson(w, req, &input)
	if err != nil {
		log.Printf("CreateEntry Input error: %v", err)

		http.Error(w, fmt.Sprintf("Error reading input: %v", err), http.StatusBadRequest)
		return
	}

	err = helpers.WriteJson(w, http.StatusOK, input, http.Header{})
	if err != nil {
		http.Error(w, "Failed to Write Json Response", http.StatusInternalServerError)
	}
}
