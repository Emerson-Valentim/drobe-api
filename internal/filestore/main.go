package filestore

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client wraps S3 operations
type Client struct {
	client *s3.Client
	bucket string
}

// NewS3Client creates a new S3 client wrapper
func NewS3Client(cfg aws.Config, bucket string, endpoint string) *Client {
	options := []func(*s3.Options){}

	if endpoint != "" {
		options = append(options, func(o *s3.Options) {
			o.UsePathStyle = true
			o.BaseEndpoint = aws.String(endpoint)
		})
	}

	s3Config := s3.NewFromConfig(cfg, options...)

	return &Client{
		client: s3Config,
		bucket: bucket,
	}
}

// GetPresignedUploadURL generates a presigned URL for uploading a file to S3
func (s *Client) GetPresignedUploadURL(ctx context.Context, reference string, expiry time.Duration) (string, string, error) {
	presignClient := s3.NewPresignClient(s.client)

	key := fmt.Sprintf("%s", reference)

	input := &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	presignedReq, err := presignClient.PresignPutObject(ctx, input, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", "", err
	}

	return key, presignedReq.URL, nil
}

// GetPresignedDownloadURL generates a presigned URL for downloading a file from S3
func (s *Client) GetPresignedDownloadURL(ctx context.Context, key string, expiry time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	presignedReq, err := presignClient.PresignGetObject(ctx, input, s3.WithPresignExpires(expiry))
	if err != nil {
		return "", err
	}

	return presignedReq.URL, nil
}

func (s *Client) DeleteObject(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	}

	_, err := s.client.DeleteObject(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
