package services

import (
	"s3-poc/data"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service struct {
	repo *data.S3Repository
}

func NewS3Service(repo *data.S3Repository) *S3Service {
	return &S3Service{repo: repo}
}

func (s *S3Service) CreateBucket(name string) error {
	return s.repo.CreateBucket(name)
}

func (s *S3Service) ListObjects(name string) ([]types.Object, error) {
	return s.repo.ListObjects(name)
}
