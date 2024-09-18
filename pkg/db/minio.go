package db

import (
	"github.com/fichca/junior-astrologer-service/internal/config"
	"github.com/minio/minio-go"
	"github.com/sirupsen/logrus"
)

func InitMinioConnection(logger *logrus.Logger, cfg *config.Minio) *minio.Client {
	minioClient, err := minio.New(cfg.Endpoint, cfg.KeyID, cfg.SecretKey, false)
	if err != nil {
		logger.Fatal(err)
	}

	ok, err := minioClient.BucketExists(cfg.Bucket)
	if err != nil {
		logger.Fatal(err)
	}

	if !ok {
		err = minioClient.MakeBucket(cfg.Bucket, "")
		if err != nil {
			logger.Fatal(err)
		}
	}

	return minioClient
}
