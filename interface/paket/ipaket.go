package ipaket

import "cukurin-barber/models"

type Repository interface {
	Create(data *models.Paket) (err error)
	Update(ID int, data map[string]interface{}) (err error)
	Delete(ID int) (err error)
}
