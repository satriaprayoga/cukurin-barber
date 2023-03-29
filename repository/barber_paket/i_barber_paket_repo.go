package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IBarberPaketRepository interface {
	GetDataBy(ID int) (result *models.BarberPaket, err error)
	GetList(queryparam models.ParamList) (result []*models.Paket, err error)
	Create(data *models.BarberPaket) error
	Update(ID int, data interface{}) error
	Delete(ID int) error
	Count(queryparam models.ParamList) (result int64, err error)
}
