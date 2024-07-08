package storage

import (
	"context"

	"github.com/thirteenths/go-minio-service/internal/models"
)

type ObjectStorage interface {
	UploadFile(context.Context, models.ObjectUnit) (string, error)
	DownloadFile(context.Context, string, string) (*models.ObjectUnit, error)
}
