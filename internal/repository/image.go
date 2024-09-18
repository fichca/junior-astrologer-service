package repository

import (
	"fmt"
	"github.com/minio/minio-go"
	"io"
	"time"
)

type Minio struct {
	minio  *minio.Client
	bucket string
}

func NewMinio(minioClient *minio.Client, bucket string) *Minio {
	return &Minio{
		minio:  minioClient,
		bucket: bucket,
	}
}

func (m *Minio) PutObject(image string, data io.Reader) error {

	_, err := m.minio.PutObject(
		m.bucket,
		image,
		data,
		-1,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return fmt.Errorf("failed to upload object to MinIO: %w", err)
	}
	return nil
}

func (m *Minio) GetImageUrls(images []string) ([]string, error) {
	var urls []string

	for i := range images {
		url, err := m.minio.PresignedGetObject(m.bucket, images[i], time.Hour*24, nil)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url.String())
	}

	return urls, nil
}

func (m *Minio) GetImageUrl(image string) (string, error) {

	url, err := m.minio.PresignedGetObject(m.bucket, image, time.Hour*24, nil)
	if err != nil {
		return "", err
	}

	return url.String(), nil
}
