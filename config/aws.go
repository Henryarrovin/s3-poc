package config

import (
	"context"
	"log"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	sdkMiddleware "github.com/aws/smithy-go/middleware"
	sdkHttp "github.com/aws/smithy-go/transport/http"
)

func NewS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider("test", "test", ""),
		),
	)
	if err != nil {
		return nil, err
	}

	endpointURL, _ := url.Parse("http://localhost:4566")

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpointURL.String())
		o.UsePathStyle = true

		o.APIOptions = append(o.APIOptions, func(stack *sdkMiddleware.Stack) error {
			stack.Serialize.Add(&s3LoggingMiddleware{}, sdkMiddleware.After)
			return nil
		})
	})

	return client, nil
}

type s3LoggingMiddleware struct{}

func (m *s3LoggingMiddleware) ID() string { return "S3Logger" }

func (m *s3LoggingMiddleware) HandleSerialize(
	ctx context.Context, in sdkMiddleware.SerializeInput, next sdkMiddleware.SerializeHandler,
) (out sdkMiddleware.SerializeOutput, metadata sdkMiddleware.Metadata, err error) {

	if req, ok := in.Request.(*sdkHttp.Request); ok {
		bucketName := "unknown"

		if opInput, ok := in.Parameters.(*s3.CreateBucketInput); ok && opInput.Bucket != nil {
			bucketName = *opInput.Bucket
		}

		log.Printf("S3 Request -> Method: %s, Bucket: %s, URL: %s\n",
			req.Method,
			bucketName,
			req.URL.String(),
		)
	} else {
		log.Println("S3 Request -> Unknown type", in.Request)
	}

	return next.HandleSerialize(ctx, in)
}
