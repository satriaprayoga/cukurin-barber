package repo

import (
	"github.com/satriaprayoga/cukurin-barber/models"
)

type IKSessionRepository interface {
	Create(data *models.KSession) error
	GetByUserID(UserID int) (output *models.KSession, err error)
	GetByAccount(account string) (output *models.KSession, err error)
	GetBySessionID(SessionID string) (output *models.KSession, err error)
	DeleteByUserID(UserID int) (err error)
	DeleteBySessionID(SessionID string) (err error)
}
