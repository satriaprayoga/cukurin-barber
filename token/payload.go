package token

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
)

type Payload struct {
	ID       uuid.UUID `json:"id"`
	UserID   int       `json:"user_id,omitempty"`
	Username string    `json:"user_name,omitempty"`
	UserType string    `json:"user_type,omitempty"`
	jwt.StandardClaims
}

func NewPayload(UserID int, Username string, UserType string) (*Payload, error) {
	tokenId := uuid.New()

	issuer := settings.AppConfigSetting.App.Issuer
	expiredTime := settings.AppConfigSetting.JWTExpired
	payload := &Payload{
		ID:       tokenId,
		Username: Username,
		UserID:   UserID,
		UserType: UserType,
		StandardClaims: jwt.StandardClaims{
			Issuer:    issuer,
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredTime)).Unix(),
		},
	}

	return payload, nil
}

// func (P *Payload) Valid() error {
// 	if time.Now().After(P.ExpiresAt) {
// 		return errors.New("expired token")
// 	}
// 	return nil
// }
