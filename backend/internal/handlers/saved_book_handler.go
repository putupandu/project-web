package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"e-library/backend/internal/models"
	"e-library/backend/internal/services"
	"e-library/backend/internal/utils"

	"github.com/gorilla/mux"
)

type SavedBookHandler struct {
	service *services.SavedBookService
}

func NewSavedBookHandler(service *services.SavedBookService) *SavedBookHandler {
	return &SavedBookHandler{service: service}
}

func (h *SavedBookHandler) SaveBook(w http.ResponseWriter, r *http.Request) {
	// Ambil user dari context
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	var req struct {
		BookID int `json:"book_id"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	if err := h.service.SaveBook(user.ID, req.BookID); err != nil {
		utils.RespondError(w, 500, err.Error())
		return
	}

	utils.RespondJSON(w, 200, models.Response{
		Success: true,
		Message: "Book saved",
	})
}

func (h *SavedBookHandler) GetSavedBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("✅ [DEBUG] Masuk ke GetSavedBooks")

	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		log.Println("❌ [ERROR] User tidak ditemukan di context")
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	log.Printf("✅ [DEBUG] User ID: %d", user.ID)

	books, err := h.service.GetSavedBooks(user.ID)
	if err != nil {
		log.Printf("❌ [ERROR] Gagal ambil saved books: %v", err)
		utils.RespondError(w, 500, err.Error())
		return
	}

	log.Printf("✅ [DEBUG] Berhasil ambil %d saved books", len(books))

	utils.RespondJSON(w, 200, models.Response{
		Success: true,
		Data:    books,
	})
}
func (h *SavedBookHandler) RemoveSavedBook(w http.ResponseWriter, r *http.Request) {
	// Ambil user dari context
	user, ok := r.Context().Value("user").(*models.User)
	if !ok || user == nil {
		utils.RespondError(w, http.StatusUnauthorized, "User not authenticated")
		return
	}

	bookID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.service.RemoveSavedBook(user.ID, bookID); err != nil {
		utils.RespondError(w, 500, err.Error())
		return
	}

	utils.RespondJSON(w, 200, models.Response{
		Success: true,
		Message: "Book removed",
	})
}
