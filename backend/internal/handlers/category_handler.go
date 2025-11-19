package handlers

import (
    "e-library/backend/internal/models"
    "e-library/backend/internal/services"
    "e-library/backend/internal/utils"
    "net/http"
    "strconv"

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