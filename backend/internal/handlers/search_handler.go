package handlers

import (
    "e-library/backend/internal/models"
    "e-library/backend/internal/services"
    "e-library/backend/internal/utils"
    "net/http"
    "strconv"
)




type SearchHandler struct {
	searchService *services.SearchService
}

func NewSearchHandler(searchService *services.SearchService) *SearchHandler {
	return &SearchHandler{searchService: searchService}
}

func (h *SearchHandler) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		utils.RespondError(w, http.StatusBadRequest, "Search query is required")
		return
	}

	filters := make(map[string]interface{})
	
	// Pagination
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	perPage := 12
	filters["page"] = page
	filters["per_page"] = perPage
	filters["limit"] = perPage
	filters["offset"] = (page - 1) * perPage

	books, meta, err := h.searchService.Search(query, filters)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Search completed successfully",
		Data:    books,
		Meta:    meta,
	})
}