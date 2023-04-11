package services

import (
	"context"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/token"
)

type IAuthService interface {
	Logout(ctx context.Context, claims token.Claims) error
	Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error)
	Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error)
	ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error)
	VerifyRegister(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error)
	VerifyRegisterLogin(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error)
	VerifyForgot(ctx context.Context, dataVerify models.VerifyForm) (err error)
}
