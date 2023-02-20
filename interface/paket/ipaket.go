package ipaket

import (
	"cukurin-barber/dto"
	"cukurin-barber/models"
)

type Repository interface {
	Create(data *models.Paket) (err error)
	Update(ID int, data map[string]interface{}) (err error)
	Delete(ID int) (err error)
	FindByID(ID int) (result *models.Paket, err error)
	Search(params dto.ParamList) (result []*models.Paket, err error)
}
