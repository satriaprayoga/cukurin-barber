package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type INotificationRepository interface {
	GetDataBy(ID int) (result *models.Notification, err error)
	GetList(queryparam models.ParamList) (result []*models.Notification, err error)
	Create(data *models.Notification) (err error)
	Update(ID int, data map[string]interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int64, err error)
}
