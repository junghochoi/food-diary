package handlers

import (
	"net/http"

	"food-diary/internal/helpers"
	"food-diary/models"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, req *http.Request) {
	// Define input struct for JSON
	var input struct {
		ID string `json:"id"`
	}

	// Parse JSON input
	if err := helpers.ReadJson(w, req, &input); err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Malformed JSON", err)
		return
	}

	// Fetch entry from database
	entry, err := models.GetEntryByID(req.Context(), h.pool, input.ID)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error retrieving entry", err)
		return
	}

	// Respond with entry data
	if err := helpers.WriteJson(w, http.StatusOK, entry, http.Header{}); err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Error writing JSON", err)
	}
}

func (h *Handlers) CreateEntry(w http.ResponseWriter, req *http.Request) {
	// Expected JSON Input
	var entry models.Entry
	// JSON Output
	var output struct {
		Success      bool         `json:"success"`
		RowsAffected uint32       `json:"rowsAffected"`
		Entry        models.Entry `json:"entry"`
	}

	// Read req.Body JSON
	err := helpers.ReadJson(w, req, &entry)
	if err != nil {
		helpers.RespondWithError(w, http.StatusBadRequest, "Malformed JSON", err)
		return
	}

	// Create Entry
	err = models.CreateEntry(req.Context(), h.pool, &entry)
	if err != nil {
		helpers.RespondWithError(w, http.StatusInternalServerError, "Could not Create Entry", err)
		return
	}

	output.Success = true
	output.RowsAffected = 1
	output.Entry = entry

	err = helpers.WriteJson(w, http.StatusOK, output, http.Header{})
	if err != nil {
		helpers.RespondWithError(
			w,
			http.StatusInternalServerError,
			"Entry created; Error writing JSON",
			err,
		)
	}
}
