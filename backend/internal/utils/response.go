package utils

import (
    "e-library/backend/internal/models"
    "encoding/json"
    "net/http"
)


func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	RespondJSON(w, status, models.ErrorResponse{
		Success: false,
		Message: message,
		Error:   message,
	})
}