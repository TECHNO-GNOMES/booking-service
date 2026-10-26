package seaweedfs

import (
	"context"
	"fmt"
	"io"
	"mime"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type FileStorage struct {
	client         *s3.Client
	bucket         string
	endpoint       string
	publicEndpoint string
}

func NewFileStorage(client *s3.Client, endpoint string, publicEndpoint string) *FileStorage {
	return &FileStorage{
		client:         client,
		bucket:         "avatars",
		endpoint:       endpoint,
		publicEndpoint: publicEndpoint,
	}
}

func (f *FileStorage) Upload(ctx context.Context, userID uint, file io.Reader, filename string) (string, error) {
	key := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), filename)

	contentType := mime.TypeByExtension(filepath.Ext(filename))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(f.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", f.publicEndpoint, f.bucket, key)

	return url, nil
}
