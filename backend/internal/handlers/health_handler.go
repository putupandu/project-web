package handlers

import (
    "e-library/backend/internal/models"
    "e-library/backend/internal/utils"
    "net/http"
)



type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Server is running healthy",
	})
}