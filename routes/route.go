package routes

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/controllers"
	"github.com/satriaprayoga/cukurin-barber/pkg/database"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	repoimpl "github.com/satriaprayoga/cukurin-barber/repository"
	"github.com/satriaprayoga/cukurin-barber/usecases"
)

type AppRoute struct {
	E *echo.Echo
}

func (a *AppRoute) Setup() {
	timeOutContext := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second

	fileRepo := repoimpl.NewRepoFileUpload(database.Conn)
	fileService := usecases.NewFileUploadSevice(fileRepo, timeOutContext)
	controllers.NewFileUploadController(a.E, fileService)
}
