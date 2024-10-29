package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
  "log"
	"food-diary/internal/helpers"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, r *http.Request) {
	mockData := map[string]string{
		"id":       "entry1",
		"mealType": "Breakfast",
		"item":     "Apple",
		"desc":     "very healthy apple",
	}

	js, err := json.Marshal(mockData)
	if err != nil {
		fmt.Errorf("error: %v", err)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")

	w.Write(js)
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


  w.Header().Set("Content-Type", "application/json")
  jsonResponse, err := json.Marshal(input)

  if err != nil {
    log.Printf("CreateEntry Output error: %v", error)
    http.Errorw(w)
  }
}
