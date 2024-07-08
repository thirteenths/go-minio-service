package miniostorage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/thirteenths/go-minio-service/internal/models"
	"github.com/thirteenths/go-minio-service/internal/storage"
)

type MinioStorage struct {
	client *minio.Client
}

func NewMinioStorage(minioURL string, minioUser string, minioPassword string, ssl bool) (storage.ObjectStorage, error) {
	var err error
	client, err := minio.New(minioURL, &minio.Options{
		Creds:  credentials.NewStaticV4(minioUser, minioPassword, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return &MinioStorage{client: client}, nil
}

// UploadFile - Отправляет файл в minio
func (m *MinioStorage) UploadFile(ctx context.Context, object models.ObjectUnit) (string, error) {
	object.ObjectName = uuid.New().String()

	_, err := m.client.PutObject(
		ctx,
		object.BucketName,
		object.ObjectName,
		bytes.NewReader(object.ObjectBytes),
		object.ObjectSize,
		minio.PutObjectOptions{ContentType: object.ContentType},
	)

	return object.ObjectName, err
}

// DownloadFile - Возвращает файл из minio
func (m *MinioStorage) DownloadFile(ctx context.Context, objectName string, bucketName string) (*models.ObjectUnit, error) {
	reader, err := m.client.GetObject(
		ctx,
		bucketName,
		objectName,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	objectInfo, err := reader.Stat()
	if err != nil {
		return nil, fmt.Errorf("failed to get file. err: %w", err)
	}

	buffer := make([]byte, objectInfo.Size)
	_, err = reader.Read(buffer)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to get objects. err: %w", err)
	}

	return &models.ObjectUnit{
		BucketName:  bucketName,
		ObjectName:  objectName,
		ObjectBytes: buffer,
		ObjectSize:  objectInfo.Size,
		ContentType: "",
	}, nil
}
