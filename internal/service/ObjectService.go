package service

import "github.com/thirteenths/go-minio-service/internal/storage"

type ObjectService struct {
	storage storage.ObjectStorage
}

func NewObjectService(storage storage.ObjectStorage) *ObjectService {
	return &ObjectService{storage: storage}
}

func (s *ObjectService) UploadObject() (string, error) {
	return s.storage.UploadFile()
}

func (s *ObjectService) DownloadFile() {
	return s.storage.DownloadFile()
}
