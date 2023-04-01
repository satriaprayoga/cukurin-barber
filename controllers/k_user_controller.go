package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/form"
	"github.com/satriaprayoga/cukurin-barber/interfaces/services"
	"github.com/satriaprayoga/cukurin-barber/middlewares"
	"github.com/satriaprayoga/cukurin-barber/pkg/response"
)

type KUserController struct {
	kuserUseCase services.IKUserService
}

func NewKUserController(e *echo.Echo, k services.IKUserService) {
	cont := &KUserController{kuserUseCase: k}
	r := e.Group("/barber/users")
	r.Use(middlewares.JWT)
	r.GET("/:id", cont.GetDataBy)

}

func (c *KUserController) GetDataBy(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp = response.Resp{R: e}
		id   = e.Param("id")
	)

	ID, err := strconv.Atoi(id)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.Response(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	data, err := c.kuserUseCase.GetDataBy(ctx, claims, ID)
	if err != nil {
		return resp.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "OK", data)
}
