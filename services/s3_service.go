package services

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Service struct {
	client *s3.Client
}

func NewS3Service(client *s3.Client) *S3Service {
	return &S3Service{client: client}
}

func (s *S3Service) CreateBucket(bucketName string) error {
	_, err := s.client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	return err
}

func (s *S3Service) ListObjects(bucketName string) ([]types.Object, error) {
	out, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, err
	}
	return out.Contents, nil
}
