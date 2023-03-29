package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IFileUploadRepository interface {
	Create(data *models.FileUpload) (err error)
	GetByID(fileID int) (models.FileUpload, error)
	Delete(fileID int) error
}
