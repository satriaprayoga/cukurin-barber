package services

import (
	"context"
	"time"

	"github.com/satriaprayoga/cukurin-barber/models"
	repo "github.com/satriaprayoga/cukurin-barber/repository"
)

type fileUploadService struct {
	repoFileUpload repo.IFileUploadRepository
	contextTimeOut time.Duration
}

func NewFileUploadSevice(a repo.IFileUploadRepository, timeout time.Duration) IFileUploadService {
	return &fileUploadService{repoFileUpload: a, contextTimeOut: timeout}
}

func (f *fileUploadService) CreateFileUpload(ctx context.Context, data *models.FileUpload) error {
	_, cancel := context.WithTimeout(ctx, f.contextTimeOut)
	defer cancel()

	var err = f.repoFileUpload.Create(data)
	if err != nil {
		return err
	}

	return err
}

func (f *fileUploadService) GetByFileID(ctx context.Context, fileID int) (models.FileUpload, error) {
	_, cancel := context.WithTimeout(ctx, f.contextTimeOut)
	defer cancel()

	var (
		err    error
		result models.FileUpload
	)

	result, err = f.repoFileUpload.GetByID(fileID)
	if err != nil {
		return result, err
	}

	return result, err
}
