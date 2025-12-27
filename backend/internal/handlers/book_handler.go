package handlers

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/services"
	"e-library/backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// 🔧 Helpers untuk membuat URL publik
func formatCoverURL(cover string) string {
	if cover == "" {
		return ""
	}
	return "http://localhost:8080/uploads/" + cover
}
func formatFileURL(file string) string {
	if file == "" {
		return ""
	}
	return "http://localhost:8080/uploads/" + file
}

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

// GET /api/books
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	filters := map[string]interface{}{}

	// search
	if s := r.URL.Query().Get("search"); s != "" {
		filters["search"] = s
	}

	// category filter
	if cid := r.URL.Query().Get("category_id"); cid != "" {
		if id, err := strconv.Atoi(cid); err == nil {
			filters["category_id"] = id
		}
	}

	// pagination
	if p := r.URL.Query().Get("page"); p != "" {
		if page, err := strconv.Atoi(p); err == nil {
			filters["page"] = page
		}
	}
	if pp := r.URL.Query().Get("per_page"); pp != "" {
		if per, err := strconv.Atoi(pp); err == nil {
			filters["per_page"] = per
		}
	}

	books, meta, err := h.bookService.GetAllBooks(filters)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Format cover & file -> public URL
	for i := range books {
		if books[i].Cover != nil {
			u := formatCoverURL(*books[i].Cover)
			books[i].Cover = &u
		}
		if books[i].FileURL != nil {
			u := formatFileURL(*books[i].FileURL)
			books[i].FileURL = &u
		}
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Books retrieved successfully",
		Data:    books,
		Meta:    meta,
	})
}

// GET /api/books/{id}
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if book == nil {
		utils.RespondError(w, http.StatusNotFound, "Book not found")
		return
	}

	// format urls
	if book.Cover != nil {
		u := formatCoverURL(*book.Cover)
		book.Cover = &u
	}
	if book.FileURL != nil {
		u := formatFileURL(*book.FileURL)
		book.FileURL = &u
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

// POST /api/books (multipart/form-data)
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// accept multipart form
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid form data: "+err.Error())
		return
	}

	book := models.Book{}
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	if y := r.FormValue("year"); y != "" {
		if yy, err := strconv.Atoi(y); err == nil {
			book.Year = yy
		}
	}
	if cid := r.FormValue("category_id"); cid != "" {
		if id, err := strconv.Atoi(cid); err == nil {
			book.CategoryID = &id
		}
	}

	// helper save file
	saveIfPresent := func(formKey string) (*string, error) {
		file, header, err := r.FormFile(formKey)
		if err != nil {
			// treat as not provided
			return nil, nil
		}
		defer file.Close()
		newName, err := utils.SaveFile("uploads", header.Filename, file)
		if err != nil {
			return nil, err
		}
		return &newName, nil
	}

	if coverName, err := saveIfPresent("cover"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to save cover: "+err.Error())
		return
	} else if coverName != nil {
		book.Cover = coverName
	}

	if fName, err := saveIfPresent("file"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to save file: "+err.Error())
		return
	} else if fName != nil {
		book.FileURL = fName
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// format before respond
	if book.Cover != nil {
		u := formatCoverURL(*book.Cover)
		book.Cover = &u
	}
	if book.FileURL != nil {
		u := formatFileURL(*book.FileURL)
		book.FileURL = &u
	}

	utils.RespondJSON(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "Book created successfully",
		Data:    book,
	})
}

// PUT /api/books/{id} (multipart/form-data allowed)
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	book := models.Book{}
	book.Title = r.FormValue("title")
	book.Author = r.FormValue("author")
	book.Description = r.FormValue("description")
	if y := r.FormValue("year"); y != "" {
		if yy, err := strconv.Atoi(y); err == nil {
			book.Year = yy
		}
	}
	if cid := r.FormValue("category_id"); cid != "" {
		if idv, err := strconv.Atoi(cid); err == nil {
			book.CategoryID = &idv
		}
	}

	saveIfPresent := func(formKey string) (*string, error) {
		file, header, err := r.FormFile(formKey)
		if err != nil {
			return nil, nil
		}
		defer file.Close()
		newName, err := utils.SaveFile("uploads", header.Filename, file)
		if err != nil {
			return nil, err
		}
		return &newName, nil
	}

	if coverName, err := saveIfPresent("cover"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to save cover: "+err.Error())
		return
	} else if coverName != nil {
		book.Cover = coverName
	}

	if fName, err := saveIfPresent("file"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to save file: "+err.Error())
		return
	} else if fName != nil {
		book.FileURL = fName
	}

	if err := h.bookService.UpdateBook(id, &book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	// format before respond
	if book.Cover != nil {
		u := formatCoverURL(*book.Cover)
		book.Cover = &u
	}
	if book.FileURL != nil {
		u := formatFileURL(*book.FileURL)
		book.FileURL = &u
	}

	book.ID = id
	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book updated successfully",
		Data:    book,
	})
}

// DELETE /api/books/{id}
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

// POST /api/books/{id}/download
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

// POST /api/books/{id}/view
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
