package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/token"
)

type IBarberService interface {
	GetDataBy(ctx context.Context, Claims token.Claims, ID int) (result interface{}, err error)
	GetDataFirstt(ctx context.Context, Claims token.Claims, ID int) (result interface{}, err error)
	GetList(ctx context.Context, Claims token.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error)
	Create(ctx context.Context, Claims token.Claims, data *models.BarbersPost) error
	Update(ctx context.Context, Claims token.Claims, ID int, data models.BarbersPost) (err error)
	Delete(ctx context.Context, Claims token.Claims, ID int) (err error)
}
