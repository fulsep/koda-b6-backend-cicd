package services

import (
	"backend/internal/models"
	"backend/internal/repository"
)

type ProfileService struct {
	userRepo *repository.UserRepo
}

func NewProfileService(userRepo *repository.UserRepo) *ProfileService {
	return &ProfileService{
		userRepo: userRepo,
	}
}

func (s *ProfileService) GetProfile(data *models.User) (*models.User, error) {
	return s.userRepo.GetUserById(&models.User{
		Id: data.Id,
	})
}

func (s *ProfileService) UpdateProfile(data *models.User) (*models.User, error) {
	return s.userRepo.UpdateUser(data)
}
