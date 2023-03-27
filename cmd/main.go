package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/satriaprayoga/cukurin-barber/pkg/database"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/sessions"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"github.com/satriaprayoga/cukurin-barber/routes"
)

func init() {
	settings.Setup("./config/config.json")
	database.Setup()
	// redisdb.Setup()
	sessions.Setup()
	logging.Setup()
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	R := routes.AppRoute{E: e}
	R.Setup()
	sPort := fmt.Sprintf(":%d", settings.AppConfigSetting.Server.HTTPPort)
	log.Fatal(e.Start(sPort))
	// settings.Setup("./config/config.json")

	// token_builder := token.NewJWTBuilder(settings.AppConfigSetting.App.JwtSecret)
	// t, _ := token_builder.CreateToken(1, utils.RandomUserName(), "user")
	// fmt.Println(t)

	// p, _ := token_builder.VerifyToken(t)
	// fmt.Printf("%v", p)

	// e := echo.New()
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	// e.Use(middleware.CORS())

	// R := routes.AppRoutes{E: e}
	// R.InitRouter()
	// sPort := fmt.Sprintf(":%d", settings.AppConfigSetting.Server.HTTPPort)
	// log.Fatal(e.Start(sPort))

	// timeOutCtx := time.Duration(settings.AppConfigSetting.Server.ReadTimeOut) * time.Second
	// repoUser := repokuser.NewRepoKUser(database.Conn)
	// useUser := usekuser.NewUseKUser(repoUser, timeOutCtx)

	// updateUser := models.UpdateUser{
	// 	UserName: "john32",
	// 	Name:     "john tulang",
	// 	Telp:     "08131222102",
	// 	Email:    "john31@gmail.com",
	// 	UserType: "editor",
	// }

	// err_ := useUser.Update(context.Background(), 1, updateUser)
	// if err_ != nil {
	// 	fmt.Println(err_)
	// } else {
	// 	fmt.Printf("OK")
	// }

}
