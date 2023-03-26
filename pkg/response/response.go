package response

import (
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
