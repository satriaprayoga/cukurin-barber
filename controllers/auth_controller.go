package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/form"
	"github.com/satriaprayoga/cukurin-barber/interfaces/services"
	"github.com/satriaprayoga/cukurin-barber/middlewares"
	"github.com/satriaprayoga/cukurin-barber/pkg/response"
)

type AuthController struct {
	authservice services.IAuthService
}

func NewAuthController(e *echo.Echo, authService services.IAuthService) {
	cont := &AuthController{
		authservice: authService,
	}
	L := e.Group("/barber/auth/logout")
	L.Use(middlewares.JWT)
	L.POST("", cont.Logout)

}

func (u *AuthController) Logout(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = response.Resp{R: e}
	)

	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	//Token := e.Request().Header.Get("Authorization")
	err = u.authservice.Logout(ctx, claims)
	if err != nil {
		return resp.ResponseError(http.StatusUnauthorized, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", nil)
}
