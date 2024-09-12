//go:generate mockgen -destination=mocks/mock_store.auto.go -package mocks github.com/anthonyvitale/moneta S3API

// Package moneta provides S3-like operations for storing files in a backing storage system.
package moneta

import (
	"context"
	"errors"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go/middleware"
)

// S3API is an interface for the functionality needed from our blob storage host.
type S3API interface {
	HeadBucket(ctx context.Context, params *s3.HeadBucketInput, optFns ...func(*s3.Options)) (*s3.HeadBucketOutput, error)
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

// S3Store describes the S3-like interface.
type S3Store interface {
	// Ping is used to check connection health
	Ping(ctx context.Context) error
	// UploadImage uploads the given content to the provided path. It returns an error indicating whether or not the
	// upload was a success.
	UploadImage(ctx context.Context, path string, body io.Reader) error
}

// Store implements the S3Store interface and is the way to interface with an S3-like product.
type Store struct {
	client  S3API
	bucket  string
	apiOpts []func(*middleware.Stack) error
}

// New creates a Store.
func New(client S3API, bucket string, apiOpts ...func(*middleware.Stack) error) (*Store, error) {
	if bucket == "" {
		return nil, errors.New("bucket name cannot be empty")
	}
	return &Store{
		client:  client,
		bucket:  bucket,
		apiOpts: apiOpts,
	}, nil
}

// Ping checks if a bucket exists. Useful as a health check.
func (s *Store) Ping(ctx context.Context) error {
	_, err := s.client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(s.bucket),
	})
	return err
}

// UploadImage uploads an object to the backing blob storage.
func (s *Store) UploadImage(ctx context.Context, key string, body io.Reader) error {
	return s.UploadImageWithMetadata(ctx, key, body, map[string]string{})
}

// UploadImageWithMetadata uploads an object with metadata to the backing blob storage.
func (s *Store) UploadImageWithMetadata(ctx context.Context, key string, body io.Reader, metadata map[string]string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if body == nil {
		return errors.New("body cannot be nil")
	}

	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:   aws.String(s.bucket),
		Key:      aws.String(key),
		Body:     body,
		Metadata: metadata,
	}, s3.WithAPIOptions(s.apiOpts...))

	return err
}
