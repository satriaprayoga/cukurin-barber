package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/token"
)

type IKUserService interface {
	GetByEmailKUser(ctx context.Context, email string, usertype string) (result models.KUser, err error)
	ChangePassword(ctx context.Context, Claims token.Claims, DataChPwd models.ChangePassword) (err error)
	GetDataBy(ctx context.Context, Claims token.Claims, ID int) (result interface{}, err error)
	GetList(ctx context.Context, Claims token.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error)
	Create(ctx context.Context, data *models.KUser) (err error)
	Update(ctx context.Context, Claims token.Claims, ID int, data models.UpdateUser) (err error)
	Delete(ctx context.Context, ID int) (err error)
}
