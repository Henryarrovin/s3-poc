package services

import (
	"s3-poc/data"
	"s3-poc/models"
)

type UserService struct {
	userRepo *data.UserRepository
}

func NewUserService(userRepo *data.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetUser(id)
}
