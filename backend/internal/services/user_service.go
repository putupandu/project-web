package services

import (
	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
)
//
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
	user := &models.User{
		Name:  name,
		Email: email,
	}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, name, email string) error {
	user := &models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return s.repo.Update(user)
}

func (s *UserService) DeleteUser(id int) error {
	return s.repo.Delete(id)
}
