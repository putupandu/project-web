// auth
package handlers

import (
	"encoding/json"
	"net/http"

	"e-library/backend/internal/models"
	"e-library/backend/internal/services"
	"e-library/backend/internal/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondError(w, 400, "invalid request")
		return
	}

	user, err := h.service.Register(req.Name, req.Email, req.Password)
	if err != nil {
		utils.RespondError(w, 400, err.Error())
		return
	}

	utils.RespondJSON(w, 201, models.Response{
		Success: true,
		Data:    user,
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	token, user, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		utils.RespondError(w, 401, err.Error())
		return
	}

	utils.RespondJSON(w, 200, models.Response{
		Success: true,
		Data: map[string]interface{}{
			"token": token,
			"user":  user,
		},
	})
}
// initial auth handler logic
