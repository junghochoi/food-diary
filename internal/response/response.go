package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`            // "success" or "error"
	Message string      `json:"message,omitempty"` // Optional: brief description of the response
	Data    interface{} `json:"data,omitempty"`    // Actual data for success responses
	Errors  interface{} `json:"errors,omitempty"`  // Error details if the request failed
}

// WriteJSON is a helper function to send JSON responses with consistent headers and status codes.
func WriteJSON(w http.ResponseWriter, status int, payload Response) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(payload)
}

// Success sends a success response with status code 200 and a data payload.
func Success(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusOK, response)
}

// Created sends a 201 Created response with data for new resources.
func Created(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusCreated, response)
}

// Error sends an error response with a specified status code and error details.
func Error(w http.ResponseWriter, status int, message string, errors interface{}) error {
	response := Response{
		Status:  "error",
		Message: message,
		Errors:  errors,
	}
	return WriteJSON(w, status, response)
}
