package handlers

import (
    "e-library/backend/internal/models"
    "e-library/backend/internal/services"
    "e-library/backend/internal/utils"
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/gorilla/mux"
)




type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	filters := make(map[string]interface{})
	
	if search := r.URL.Query().Get("search"); search != "" {
		filters["search"] = search
	}
	
	if categoryID := r.URL.Query().Get("category_id"); categoryID != "" {
		if id, err := strconv.Atoi(categoryID); err == nil {
			filters["category_id"] = id
		}
	}
	
	if year := r.URL.Query().Get("year"); year != "" {
		if y, err := strconv.Atoi(year); err == nil {
			filters["year"] = y
		}
	}

	// Pagination
	page := 1
	if p := r.URL.Query().Get("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}
	filters["page"] = page

	perPage := 12
	if pp := r.URL.Query().Get("per_page"); pp != "" {
		if parsed, err := strconv.Atoi(pp); err == nil && parsed > 0 {
			perPage = parsed
		}
	}
	filters["per_page"] = perPage
	filters["limit"] = perPage
	filters["offset"] = (page - 1) * perPage

	// Sorting
	if sort := r.URL.Query().Get("sort"); sort != "" {
		filters["sort"] = sort
	}
	if order := r.URL.Query().Get("order"); order != "" {
		filters["order"] = order
	}

	books, meta, err := h.bookService.GetAllBooks(filters)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Books retrieved successfully",
		Data:    books,
		Meta:    meta,
	})
}

func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "Book created successfully",
		Data:    book,
	})
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.bookService.UpdateBook(id, &book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	book.ID = id
	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book updated successfully",
		Data:    book,
	})
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.bookService.DeleteBook(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book deleted successfully",
	})
}

func (h *BookHandler) IncrementDownload(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.bookService.IncrementDownload(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Download count incremented",
	})
}

func (h *BookHandler) IncrementView(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.bookService.IncrementView(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "View count incremented",
	})
}