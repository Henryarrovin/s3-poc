package data

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Repository struct {
	client *s3.Client
}

func NewS3Repository(client *s3.Client) *S3Repository {
	return &S3Repository{client: client}
}

func (r *S3Repository) CreateBucket(bucketName string) error {
	_, err := r.client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	return err
}

func (r *S3Repository) ListObjects(bucketName string) ([]types.Object, error) {
	out, err := r.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: &bucketName,
	})
	if err != nil {
		return nil, err
	}
	return out.Contents, nil
}
