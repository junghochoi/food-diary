package handlers

import (
	"net/http"

	"food-diary/internal/helpers"
	"food-diary/models"
)

func (h *Handlers) GetEntry(w http.ResponseWriter, req *http.Request) {
	// Expected Input
	var input struct {
		Id string `json:"id"`
	}

	// Read req.Body JSON into input variable
	err := helpers.ReadJson(w, req, &input)
	if err != nil {
		helpers.RespondWithError(
			w,
			http.StatusBadRequest,
			"Malformed JSON",
			err,
		)
		return
	}

	// Fetch Entry Model of ID from database
	entry, err := models.GetEntryByID(req.Context(), h.pool, input.Id)
	if err != nil {
		helpers.RespondWithError(
			w,
			http.StatusInternalServerError,
			"Error Retrieving Entry",
			err,
		)
		return

	}

	// Write Result to Response Body
	err = helpers.WriteJson(w, http.StatusOK, entry, http.Header{})
	if err != nil {
		helpers.RespondWithError(
			w,
			http.StatusBadRequest,
			"Error Writing JSON",
			err,
		)
		return
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
