package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-barber/models"
)

type IFileUploadService interface {
	CreateFileUpload(ctx context.Context, data *models.FileUpload) error
	GetByFileID(ctx context.Context, fileID int) (models.FileUpload, error)
}
