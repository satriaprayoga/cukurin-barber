package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type ICapsterCollectionRepository interface {
	GetDataBy(ID int) (result *models.CapsterCollection, err error)
	GetListFileCapter(ID int) (result []*models.FileOutput, err error)
	GetList(queryparam models.ParamList) (result []*models.CapsterList, err error)
	Create(data *models.CapsterCollection) error
	Update(ID int, data interface{}) error
	Delete(ID int) error
	Count(queryparam models.ParamList) (result int64, err error)
}
