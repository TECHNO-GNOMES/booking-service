package seaweedfs

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type FileStorage struct {
	client   *s3.Client
	bucket   string
	endpoint string
}

func NewFileStorage(client *s3.Client, endpoint string) *FileStorage {
	return &FileStorage{
		client:   client,
		bucket:   "avatars",
		endpoint: endpoint,
	}
}

func (f *FileStorage) Upload(ctx context.Context, userID uint, file io.Reader, filename string) (string, error) {
	key := fmt.Sprintf("%d_%d_%s", userID, time.Now().Unix(), filename)

	_, err := f.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(f.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s/%s/%s", f.endpoint, f.bucket, key)

	return url, nil
}
