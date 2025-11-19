package services

import (
    "errors"
    "e-library/backend/internal/models"
    "e-library/backend/internal/repositories"
)



type BookService struct {
	bookRepo     *repositories.BookRepository
	categoryRepo *repositories.CategoryRepository
}

func NewBookService(bookRepo *repositories.BookRepository, categoryRepo *repositories.CategoryRepository) *BookService {
	return &BookService{
		bookRepo:     bookRepo,
		categoryRepo: categoryRepo,
	}
}

func (s *BookService) GetAllBooks(filters map[string]interface{}) ([]models.Book, *models.Meta, error) {
	books, err := s.bookRepo.FindAll(filters)
	if err != nil {
		return nil, nil, err
	}

	// Get total count for pagination
	total, err := s.bookRepo.Count(filters)
	if err != nil {
		return books, nil, nil
	}

	// Calculate pagination meta
	page := 1
	perPage := 12
	if p, ok := filters["page"].(int); ok && p > 0 {
		page = p
	}
	if pp, ok := filters["per_page"].(int); ok && pp > 0 {
		perPage = pp
	}

	totalPages := (total + perPage - 1) / perPage

	meta := &models.Meta{
		Total:      total,
		Page:       page,
		PerPage:    perPage,
		TotalPages: totalPages,
	}

	return books, meta, nil
}

func (s *BookService) GetBookByID(id int) (*models.Book, error) {
	book, err := s.bookRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	return book, nil
}

func (s *BookService) CreateBook(book *models.Book) error {
	// Validate category if provided
	if book.CategoryID != nil {
		cat, err := s.categoryRepo.FindByID(*book.CategoryID)
		if err != nil {
			return err
		}
		if cat == nil {
			return errors.New("category not found")
		}
	}

	return s.bookRepo.Create(book)
}

func (s *BookService) UpdateBook(id int, book *models.Book) error {
	// Check if book exists
	existing, err := s.bookRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return errors.New("book not found")
	}

	// Validate category if provided
	if book.CategoryID != nil {
		cat, err := s.categoryRepo.FindByID(*book.CategoryID)
		if err != nil {
			return err
		}
		if cat == nil {
			return errors.New("category not found")
		}
	}

	return s.bookRepo.Update(id, book)
}

func (s *BookService) DeleteBook(id int) error {
	return s.bookRepo.Delete(id)
}

func (s *BookService) IncrementDownload(id int) error {
	return s.bookRepo.IncrementDownloads(id)
}

func (s *BookService) IncrementView(id int) error {
	return s.bookRepo.IncrementViews(id)
}