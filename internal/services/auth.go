package services

import (
	"backend/internal/models"
	"backend/internal/repository"
	"errors"
)

type AuthService struct {
	userRepo *repository.UserRepo
}

func NewAuthService(userRepo *repository.UserRepo) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Login(data *models.User) (*models.User, error) {
	user, err := s.userRepo.GetUserByEmail(data)

	if err != nil {
		return nil, err
	}

	if data.Password != user.Password {
		return nil, errors.New("Wrong password")
	}

	return user, nil
}
