package services

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
)

type SearchService struct {
	bookRepo *repositories.BookRepository
}
//
func NewSearchService(bookRepo *repositories.BookRepository) *SearchService {
	return &SearchService{bookRepo: bookRepo}
}

func (s *SearchService) Search(query string, filters map[string]interface{}) ([]models.Book, *models.Meta, error) {
	filters["search"] = query
	
	books, err := s.bookRepo.FindAll(filters)
	if err != nil {
		return nil, nil, err
	}

	total, _ := s.bookRepo.Count(filters)
	
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