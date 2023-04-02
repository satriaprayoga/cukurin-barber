package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/satriaprayoga/cukurin-barber/form"
	"github.com/satriaprayoga/cukurin-barber/interfaces/services"
	"github.com/satriaprayoga/cukurin-barber/middlewares"
	"github.com/satriaprayoga/cukurin-barber/models"
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
	r.GET("", cont.GetList)
	r.POST("", cont.Create)
	r.POST("/change_password", cont.ChangePassword)
	r.PUT("/:id", cont.Update)
	r.DELETE("/:id", cont.Delete)

}

func (u *KUserController) GetDataBy(e echo.Context) error {
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
	data, err := u.kuserUseCase.GetDataBy(ctx, claims, ID)
	if err != nil {
		return resp.Response(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "OK", data)
}

func (u *KUserController) GetList(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp         = response.Resp{R: e}
		paramQuery   = models.ParamList{}
		responseList = models.ResponseModelList{}
		err          error
	)
	httpCode, errMsg := form.BindAndValid(e, &paramQuery)
	if httpCode != 200 {
		return resp.ResponseErrorList(http.StatusBadRequest, errMsg, responseList)
	}
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.ResponseErrorList(http.StatusBadRequest, fmt.Sprintf("%v", err), responseList)
	}

	responseList, err = u.kuserUseCase.GetList(ctx, claims, paramQuery)
	if err != nil {
		return resp.ResponseErrorList(response.GetStatusCode(err), fmt.Sprintf("%v", err), responseList)
	}
	return resp.Response(http.StatusOK, "", responseList)
}

func (u *KUserController) Create(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	var (
		resp     = response.Resp{R: e}
		kuser    models.KUser
		userForm models.AddUser
	)

	httpCode, errMsg := form.BindAndValid(e, &userForm)
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err := mapstructure.Decode(userForm, &kuser)
	if err != nil {
		return resp.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	err = u.kuserUseCase.Create(ctx, claims, &kuser)
	if err != nil {
		return resp.ResponseError(response.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusCreated, "Ok", kuser)
}

func (u *KUserController) Update(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp     = response.Resp{R: e}
		err      error
		id       = e.Param("id")
		userForm = models.UpdateUser{}
	)

	OwnerID, err := strconv.Atoi(id)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	httpCode, errMsg := form.BindAndValid(e, &userForm)
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err = u.kuserUseCase.Update(ctx, claims, OwnerID, userForm)
	if err != nil {
		return resp.ResponseError(response.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusCreated, "Ok", nil)

}

func (u *KUserController) ChangePassword(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	var (
		resp     = response.Resp{R: e}
		err      error
		userForm = models.ChangePassword{}
	)
	claims, err := form.GetClaims(e)
	if err != nil {
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	httpCode, errMsg := form.BindAndValid(e, &userForm)
	if httpCode != 200 {
		return resp.ResponseError(http.StatusBadRequest, errMsg, nil)
	}
	err = u.kuserUseCase.ChangePassword(ctx, claims, userForm)
	if err != nil {
		return resp.ResponseError(response.GetStatusCode(err), fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusCreated, "Ok", nil)

}

func (u *KUserController) Delete(e echo.Context) error {
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
		return resp.ResponseError(http.StatusBadRequest, fmt.Sprintf("%v", err), nil)
	}
	err = u.kuserUseCase.Delete(ctx, ID)
	if err != nil {
		return resp.ResponseError(http.StatusInternalServerError, fmt.Sprintf("%v", err), nil)
	}
	return resp.Response(http.StatusOK, "Ok", nil)

}
