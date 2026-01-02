package services

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) CreateUser(name, email string) (*models.User, error) {
	return s.repo.Create(name, email)
}

func (s *UserService) UpdateUser(id int, name, email string) error {
	return s.repo.Update(id, name, email)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}