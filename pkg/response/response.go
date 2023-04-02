package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/utils"
)

type Resp struct {
	R echo.Context
}

type ResponseModel struct {
	Message string      `jsong:"message"`
	Data    interface{} `json:"data"`
}

func (e Resp) Response(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{
		Message: errMsg,
		Data:    data,
	}
	logger.Info(string(utils.Stringify(response)))
	return e.R.JSON(httpCode, response)
}

func (e Resp) ResponseError(httpCode int, errMsg string, data interface{}) error {
	var logger = logging.Logger{}
	response := ResponseModel{

		Message: errMsg,
		Data:    data,
	}
	logger.Error(string(utils.Stringify(response)))
	return e.R.JSON(httpCode, response)
	// return string(util.Stringify(response))
}

// ResponseErrorList :
func (e Resp) ResponseErrorList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Error(string(utils.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}

// ResponseList :
func (e Resp) ResponseList(httpCode int, errMsg string, data models.ResponseModelList) error {
	var logger = logging.Logger{}
	data.Msg = errMsg

	logger.Info(string(utils.Stringify(data)))
	return e.R.JSON(httpCode, data)
	// return string(util.Stringify(response))
}

func GetStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	switch err {
	case models.ErrInternalServerError:
		return http.StatusInternalServerError
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrConflict:
		return http.StatusConflict
	case models.Unauthorized:
		return http.StatusUnauthorized
	case models.ErrInvalidLogin:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}
