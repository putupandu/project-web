//
package handlers

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/services"
	"e-library/backend/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ================= URL DINAMIS (AMAN SAAT IP BERUBAH) =================
func makePublicURL(r *http.Request, file string) string {
	if file == "" {
		return ""
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	return scheme + "://" + r.Host + "/uploads/" + file
}

type BookHandler struct {
	bookService *services.BookService
}

func NewBookHandler(bookService *services.BookService) *BookHandler {
	return &BookHandler{bookService: bookService}
}

// ================= GET ALL BOOKS =================
func (h *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	filters := map[string]interface{}{}

	if s := r.URL.Query().Get("search"); s != "" {
		filters["search"] = s
	}

	if cid := r.URL.Query().Get("category_id"); cid != "" {
		if id, err := strconv.Atoi(cid); err == nil {
			filters["category_id"] = id
		}
	}

	books, meta, err := h.bookService.GetAllBooks(filters)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range books {

		if books[i].Language == "" {
		books[i].Language = "-"
		}
		if books[i].Cover != nil {
			url := makePublicURL(r, *books[i].Cover)
			books[i].Cover = &url
		}
		if books[i].FileURL != nil {
			url := makePublicURL(r, *books[i].FileURL)
			books[i].FileURL = &url
		}
	}
	

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Books retrieved successfully",
		Data:    books,
		Meta:    meta,
	})
}

// ================= GET BOOK BY ID =================
func (h *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
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

	if book.Cover != nil {
		url := makePublicURL(r, *book.Cover)
		book.Cover = &url
	}
	if book.FileURL != nil {
		url := makePublicURL(r, *book.FileURL)
		book.FileURL = &url
	}
	if book.Language == "" {
	book.Language = "-"
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

// ================= CREATE BOOK =================
func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	book := models.Book{
	Title:       r.FormValue("title"),
	Author:      r.FormValue("author"),
	Description: r.FormValue("description"),
	Language:    r.FormValue("language"),
}


	if y := r.FormValue("year"); y != "" {
		if year, err := strconv.Atoi(y); err == nil {
			book.Year = year
		}
	}

	if cid := r.FormValue("category_id"); cid != "" {
		if id, err := strconv.Atoi(cid); err == nil {
			book.CategoryID = &id
		}
	}

	saveUpload := func(key string) (*string, error) {
		file, header, err := r.FormFile(key)
		if err != nil {
			return nil, nil
		}
		defer file.Close()

		name, err := utils.SaveFile("uploads", header.Filename, file)
		if err != nil {
			return nil, err
		}
		return &name, nil
	}

	if c, err := saveUpload("cover"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else if c != nil {
		book.Cover = c
	}

	if f, err := saveUpload("file"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else if f != nil {
		book.FileURL = f
	}

	if err := h.bookService.CreateBook(&book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if book.Cover != nil {
		url := makePublicURL(r, *book.Cover)
		book.Cover = &url
	}
	if book.FileURL != nil {
		url := makePublicURL(r, *book.FileURL)
		book.FileURL = &url
	}

	utils.RespondJSON(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "Book created successfully",
		Data:    book,
	})
}

// ================= UPDATE BOOK =================
func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Failed to parse form")
		return
	}

	book := models.Book{
	Title:       r.FormValue("title"),
	Author:      r.FormValue("author"),
	Description: r.FormValue("description"),
	Language:    r.FormValue("language"),
	}


	if y := r.FormValue("year"); y != "" {
		if year, err := strconv.Atoi(y); err == nil {
			book.Year = year
		}
	}

	if cid := r.FormValue("category_id"); cid != "" {
		if cidInt, err := strconv.Atoi(cid); err == nil {
			book.CategoryID = &cidInt
		}
	}

	saveUpload := func(key string) (*string, error) {
		file, header, err := r.FormFile(key)
		if err != nil {
			return nil, nil
		}
		defer file.Close()

		name, err := utils.SaveFile("uploads", header.Filename, file)
		if err != nil {
			return nil, err
		}
		return &name, nil
	}

	if c, err := saveUpload("cover"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else if c != nil {
		book.Cover = c
	}

	if f, err := saveUpload("file"); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	} else if f != nil {
		book.FileURL = f
	}

	if err := h.bookService.UpdateBook(id, &book); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	book.ID = id

	if book.Cover != nil {
		url := makePublicURL(r, *book.Cover)
		book.Cover = &url
	}
	if book.FileURL != nil {
		url := makePublicURL(r, *book.FileURL)
		book.FileURL = &url
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Book updated successfully",
		Data:    book,
	})
}

// ================= DELETE BOOK =================
func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
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

// ================= INCREMENT DOWNLOAD =================
func (h *BookHandler) IncrementDownload(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.bookService.IncrementDownload(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to increment download count")
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Download count incremented",
	})
}

// ================= INCREMENT VIEW =================
func (h *BookHandler) IncrementView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	if err := h.bookService.IncrementView(id); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to increment view count")
		return
	}

	utils.RespondJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "View count incremented",
	})
}
