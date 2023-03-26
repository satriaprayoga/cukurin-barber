package token

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type JWTBuilder struct {
	secret string
}

func NewJWTBuilder(secret string) TokenBuilder {
	return &JWTBuilder{secret: secret}
}

func (builder *JWTBuilder) CreateToken(UserID int, Username string, UserType string) (string, error) {
	payload, err := NewPayload(UserID, Username, UserType)
	if err != nil {
		return "", err
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(builder.secret))
}

func (builder *JWTBuilder) VerifyToken(token string) (*Payload, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(builder.secret), nil
	})
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errors.New("expired token")) {
			return nil, errors.New("expired token")
		}
	}
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errors.New("invalid token")
	}
	return payload, nil
}
