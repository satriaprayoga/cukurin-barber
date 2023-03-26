package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
)

type Claims struct {
	ID       string `json:"id"`
	UserID   int    `json:"user_id,omitempty"`
	Username string `json:"user_name,omitempty"`
	UserType string `json:"user_type,omitempty"`
	jwt.StandardClaims
}

func GenerateToken(ID string, UserID int, Username string, UserType string) (string, error) {
	var (
		secret      = settings.AppConfigSetting.App.JwtSecret
		issuer      = settings.AppConfigSetting.App.Issuer
		expiredTime = settings.AppConfigSetting.JWTExpired
	)
	var jwt_secret = []byte(secret)

	claims := &Claims{
		ID:       ID,
		UserID:   UserID,
		Username: Username,
		UserType: UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix(),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwt_secret)
}

func ParseToken(token string) (*Claims, error) {
	var (
		secret = settings.AppConfigSetting.App.JwtSecret
	)
	var jwt_secret = []byte(secret)

	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwt_secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
