package uploader

import (
	"context"
	"mime/multipart"
)

type FileUploaderInterface interface {
	UploadFile(ctx context.Context, bucketName string, fileHeader *multipart.FileHeader) (string, error)
}
