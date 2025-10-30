package uploader

import (
	"ams-sentuh/config"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"mime/multipart"
	"time"
)

type minioUploader struct {
	cfg *config.Config
}

func NewMinioUploader(cfg *config.Config) FileUploaderInterface {
	return &minioUploader{cfg: cfg}
}

func (m *minioUploader) UploadFile(ctx context.Context, bucketName string, fileHeader *multipart.FileHeader) (string, error) {
	minioClient, err := NewMinioClient(m.cfg)
	if err != nil {
		return "", errors.Wrap(err, "failed to init minio client")
	}

	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", errors.Wrap(err, "failed to check bucket")
	}
	if !exists {
		if err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{}); err != nil {
			return "", errors.Wrap(err, "failed to create bucket")
		}
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer file.Close()

	objectName := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)
	_, err = minioClient.PutObject(ctx, bucketName, objectName, file, fileHeader.Size, minio.PutObjectOptions{
		ContentType: fileHeader.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", errors.Wrap(err, "failed to upload file to minio")
	}

	return fmt.Sprintf("http://%s/%s/%s", m.cfg.Minio.MinioEndpoint, bucketName, objectName), nil
}
