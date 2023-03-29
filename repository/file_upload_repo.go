package repoimpl

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"gorm.io/gorm"
)

type repoFileUpload struct {
	Conn *gorm.DB
}

func NewRepoFileUpload(Conn *gorm.DB) repo.IFileUploadRepository {
	return &repoFileUpload{Conn}
}

func (r *repoFileUpload) Create(data *models.FileUpload) (err error) {
	var logger = logging.Logger{}
	query := r.Conn.Create(&data)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repoFileUpload) GetByID(fileID int) (models.FileUpload, error) {
	var (
		dataFileUpload = models.FileUpload{}
		logger         = logging.Logger{}
		err            error
	)
	query := r.Conn.Where("file_id=?", fileID).First(&dataFileUpload)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dataFileUpload, models.ErrNotFound
		}
	}
	return dataFileUpload, err
}

func (r *repoFileUpload) Delete(fileID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)

	fileData := &models.FileUpload{}
	fileData.FileID = fileID

	query := r.Conn.Delete(&fileData)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil

}
