package uploader

import (
	"ams-sentuh/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/pkg/errors"
)

func NewMinioClient(cfg *config.Config) (*minio.Client, error) {
	client, err := minio.New(cfg.Minio.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		return nil, errors.Wrapf(err, "Failed to create minio client")
	}
	return client, nil
}
