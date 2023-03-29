package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IBarberCapsterRepository interface {
	GetDataBy(ID int) (result *models.BarberCapster, err error)
	GetList(queryparam models.ParamList) (result []*models.CapsterList, err error)
	Create(data *models.BarberCapster) error
	Update(ID int, data interface{}) error
	Delete(ID int) error
	DeleteByCapster(ID int) error
}
