package services

import (
	"errors"
	"time"

	"e-library/backend/internal/models"
	"e-library/backend/internal/repositories"
	"e-library/backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

//  REGISTER 
func (s *AuthService) Register(name, email, password string) (*models.User, error) {
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

//  LOGIN 
func (s *AuthService) Login(email, password string) (string, *models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, time.Hour*24)
	if err != nil {
		return "", nil, err
	}

	user.Password = ""
	return token, user, nil
}
