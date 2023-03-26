package form

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/satriaprayoga/cukurin-barber/token"
)

func BindAndValid(c echo.Context, form interface{}) (int, string) {
	err := c.Bind(form)
	if err != nil {
		return http.StatusBadRequest, "invalid request parameter"
	}
	valid := validation.Validation{}
	check, err := valid.Valid(form)
	if err != nil {
		return http.StatusInternalServerError, "internal server error"
	}
	if !check {
		return http.StatusBadRequest, MarkErrors(valid.Errors)
	}
	return http.StatusOK, "ok"
}

func MarkErrors(errors []*validation.Error) string {
	res := ""
	for _, err := range errors {
		res = fmt.Sprintf("%s %s", err.Key, err.Message)
	}
	return res
}

func GetClaims(c echo.Context) (token.Claims, error) {
	var clm token.Claims
	claims := c.Get("claims")
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(*util.Claims)

	err := mapstructure.Decode(claims, &clm)
	if err != nil {
		return clm, err
	}
	return clm, nil
}
