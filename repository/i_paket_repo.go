package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IPaketRepository interface {
	GetDataBy(ID int) (result *models.Paket, err error)
	GetList(queryparam models.ParamList) (result []*models.Paket, err error)
	Create(data *models.Paket) (err error)
	Update(ID int, data map[string]interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int64, err error)
}
