package services

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
)

type SavedBookService struct {
	repo *repositories.SavedBookRepository
}
//
func NewSavedBookService(r *repositories.SavedBookRepository) *SavedBookService {
	return &SavedBookService{repo: r}
}

func (s *SavedBookService) SaveBook(userID, bookID int) error {
	return s.repo.Save(userID, bookID)
}

func (s *SavedBookService) GetSavedBooks(userID int) ([]models.Book, error) {
	return s.repo.FindByUser(userID)
}

func (s *SavedBookService) RemoveSavedBook(userID, bookID int) error {
	return s.repo.Delete(userID, bookID)
}
