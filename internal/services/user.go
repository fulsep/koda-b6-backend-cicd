package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(data *models.User) (*models.User, error) {
	user, err := s.repo.CreateUser(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() (*[]models.User, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserByEmail(data *models.User) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserById(data *models.User) (*models.User, error) {
	user, err := s.repo.GetUserById(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(data *models.User) (*models.User, error) {
	user, err := s.repo.UpdateUser(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(data *models.User) (*models.User, error) {
	user, err := s.repo.DeleteUser(data)
	if err != nil {
		return nil, err
	}
	return user, nil
}
