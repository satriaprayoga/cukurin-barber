package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/satriaprayoga/cukurin-barber/pkg/response"
	"github.com/satriaprayoga/cukurin-barber/pkg/sessions"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"github.com/satriaprayoga/cukurin-barber/token"
)

func JWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var (
			code = http.StatusOK
			msg  = ""
			data interface{}
			t    = c.Request().Header.Get("Authorization")
		)

		data = map[string]string{
			"token": t,
		}

		if t == "" {
			code = http.StatusNetworkAuthenticationRequired
			msg = "Auth Token Required"
		} else {
			claims, err := token.ParseToken(t)
			if err != nil {
				code = http.StatusUnauthorized
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					msg = "Token Expired"
				default:
					msg = "Token Failed"
				}
			}
			_, errS := sessions.GetSession(claims.ID)
			if errS != nil {
				code = http.StatusUnauthorized
				msg = "session not valid"
			} else {
				var issuer = settings.AppConfigSetting.App.Issuer
				valid := claims.VerifyIssuer(issuer, true)
				if !valid {
					code = http.StatusUnauthorized
					msg = "issuer not found"
				}
				c.Set("claims", claims)
			}
		}
		if code != http.StatusOK {
			resp := response.ResponseModel{
				Message: msg,
				Data:    data,
			}
			return c.JSON(code, resp)
		}
		return next(c)

	}
}
