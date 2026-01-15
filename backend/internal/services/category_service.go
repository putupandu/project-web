package services

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
	"errors"
)
//
type CategoryService struct {
	categoryRepo *repositories.CategoryRepository
}

func NewCategoryService(categoryRepo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}

func (s *CategoryService) GetAllCategories() ([]models.Category, error) {
	return s.categoryRepo.FindAll()
}

func (s *CategoryService) GetCategoryByID(id int) (*models.Category, error) {
	cat, err := s.categoryRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return cat, nil
}

func (s *CategoryService) GetCategoryBySlug(slug string) (*models.Category, error) {
	cat, err := s.categoryRepo.FindBySlug(slug)
	if err != nil {
		return nil, err
	}
	if cat == nil {
		return nil, errors.New("category not found")
	}
	return cat, nil
}

func (s *CategoryService) DeleteCategory(id int) error {
	return s.categoryRepo.Delete(id)
}

func (s *CategoryService) CreateCategory(cat *models.Category) (*models.Category, error) {
	return s.categoryRepo.Create(cat)
}