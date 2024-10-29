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
		Foods       []string `json:"items"`
		Description string   `json:"desc"`
	}

	var output struct {
		Success bool `json:"success"`
		Entry   struct {
			MealType    string   `json:"mealType"`
			Foods       []string `json:"items"`
			Description string   `json:"desc"`
    } `json:"entry"`
	}

	err := helpers.ReadJson(w, req, &input)

	if err != nil {
		log.Printf("CreateEntry Input error: %v", err)

		http.Error(w, fmt.Sprintf("Error reading input: %v", err), http.StatusBadRequest)
		return
	}

	output.Success = true
	output.Entry.MealType = input.MealType
	output.Entry.Foods = input.Foods
	output.Entry.Description = input.Description


	err = helpers.WriteJson(w, http.StatusOK, output, http.Header{})
	if err != nil {
		http.Error(w, "Failed to Write Json Response", http.StatusInternalServerError)
	}
}
