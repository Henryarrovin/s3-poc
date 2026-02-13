package services

import (
	"s3-poc/data"
	"s3-poc/models"

	"gorm.io/gorm"
)

type UserService struct {
	userRepo *data.UserRepository
	s3Repo   *data.S3Repository
}

func NewUserService(
	userRepo *data.UserRepository,
	s3Repo *data.S3Repository,
) *UserService {
	return &UserService{
		userRepo: userRepo,
		s3Repo:   s3Repo,
	}
}

func (s *UserService) CreateUser(user *models.User) error {
	db := s.userRepo.DB()
	return db.Transaction(func(tx *gorm.DB) error {
		if err := s.userRepo.CreateUserTx(tx, user); err != nil {
			return err
		}

		if err := s.s3Repo.CreateBucket(user.BucketName); err != nil {
			return err
		}
		return nil
	})
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetUser(id)
}
