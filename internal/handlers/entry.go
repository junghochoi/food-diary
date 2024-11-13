package handlers

import (
	"net/http"

	"food-diary/internal/request"
	"food-diary/internal/response"
  "food-diary/internal/request/json"
	"food-diary/models"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, req *http.Request) {
	// Define input struct for JSON
	var entryGetRequest request.EntryGetRequest
	// Parse JSON input
	if err := json.DecodeRequestBody(w, req, &entryGetRequest); err != nil {
    response.Error(w,http.StatusBadRequest, "Malformed JSON", err)
		return
	}
	// TODO Add Request Input Validation

	// Fetch entry from database
	entry, err := models.GetEntryByID(req.Context(), h.pool, entryGetRequest.ID)
	if err != nil {
    response.Error(w, http.StatusBadRequest, "Malformed JSON", err)
		return
	}

	// Respond with entry data
	if err := response.Success(w, "success", &entry); err != nil {
    response.Error(w, http.StatusInternalServerError, "Error writing JSON", err )
	}
}

func (h *Handlers) CreateEntry(w http.ResponseWriter, req *http.Request) {
	// Expected JSON Input
	var entryCreateRequest request.EntryCreateRequest

	// Read req.Body JSON
	if err := json.DecodeRequestBody(w, req, &entryCreateRequest); err != nil {
		response.Error(w, http.StatusBadRequest, "Malformed JSON", err)
		return
	}

	// TODO Add Request Input Validation

	// Create Entry Instance
	entry := models.Entry{
		Title:      entryCreateRequest.Title,
		Foods:      entryCreateRequest.Foods,
		FoodDesc:   entryCreateRequest.FoodDesc,
		Rating:     entryCreateRequest.Rating,
		RatingDesc: entryCreateRequest.RatingDesc,
	}

	// Insert Entry to DB
	if err := models.CreateEntry(req.Context(), h.pool, &entry); err != nil {
		response.Error(w, http.StatusInternalServerError, "Could not Create Entry", err)
		return
	}

	// Return a response
	if err := response.Created(w, "Entry Successfully Created", &entry); err != nil {
		response.Error(w, http.StatusInternalServerError, "Failed to write reponse", err)
    return
	}
}
