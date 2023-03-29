package repo

import "github.com/satriaprayoga/cukurin-barber/models"

type IKUserRepository interface {
	GenUserCapster() (string, error)
	GetDataBy(ID int) (result *models.KUser, err error)
	GetByAccount(Account string, userType string /* , IsOwner bool */) (result models.KUser, err error)
	GetByCapster(Account string) (result models.LoginCapster, err error)
	GetList(queryparam models.ParamList) (result []*models.UserList, err error)
	Create(data *models.KUser) (err error)
	UpdatePasswordByEmail(Email string, Password string) error
	Update(ID int, data interface{}) (err error)
	Delete(ID int) (err error)
	Count(queryparam models.ParamList) (result int, err error)
}
