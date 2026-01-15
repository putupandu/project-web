package handlers

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/services"
	"e-library/backend/internal/utils"
	"log"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"

	"github.com/gorilla/mux"
)




type CategoryHandler struct {
	categoryService *services.CategoryService
}

func NewCategoryHandler(categoryService *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Categories retrieved successfully",
		Data:    categories,
	})
}

func (h *CategoryHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid category ID")
		return
	}

	category, err := h.categoryService.GetCategoryByID(id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Category retrieved successfully",
		Data:    category,
	})
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := h.categoryService.DeleteCategory(id); err != nil {
		log.Printf("Delete category error: %v", err)
		http.Error(w, "Failed to delete category", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Slug        string `json:"slug,omitempty"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Jika slug kosong, buat dari name
	slug := req.Slug
	if slug == "" {
		slug = strings.ToLower(strings.ReplaceAll(req.Name, " ", "-"))
	}

	category := &models.Category{
		Name:        req.Name,
		Description: req.Description,
		Slug:        slug,
	}

	newCat, err := h.categoryService.CreateCategory(category)
	if err != nil {
		log.Printf("Create category error: %v", err)
		http.Error(w, "Failed to create category", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCat)
}
//