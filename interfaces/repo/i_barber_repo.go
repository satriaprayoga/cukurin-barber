package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IBarberRepository interface {
	GetDataBy(ID int) (result *models.Barber, err error)
	GetDataFirst(OwnerID int, BarberID int) (result *models.Barber, err error)
	GetList(queryparam models.ParamList) (result []*models.BarbersList, err error)
	Create(data *models.Barber) (err error)
	Update(ID int, data interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int64, err error)
}
