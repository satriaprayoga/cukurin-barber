package sessions

import (
	"errors"
	"log"
	"time"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/database"
	repo "github.com/satriaprayoga/cukurin-barber/repository/k_session"
)

var sessionRepo repo.IKSessionRepository

func Setup() {
	now := time.Now()
	sessionRepo = repo.NewRepoKSession(database.Conn)
	timeSpent := time.Since(now)
	log.Printf("Config session is ready in %v", timeSpent)
}

func GetSession(sessionID string) (*models.KSession, error) {
	k_session, err := sessionRepo.GetBySessionID(sessionID)
	if err != nil {
		return nil, errors.New("session not found")
	}
	return k_session, nil
}

func GetSessionByAccount(Account string) (*models.KSession, error) {
	k_session, err := sessionRepo.GetByAccount(Account)
	if err != nil {
		return nil, errors.New("session not found")
	}
	return k_session, nil
}

func CreateSession(sessionID, sessionType, account string, userId int, expired time.Time) error {
	var ksession models.KSession
	ksession.SessionID = sessionID
	ksession.Account = account
	ksession.SessionType = sessionType
	ksession.UserID = userId
	ksession.ExpiresAt = expired
	err := sessionRepo.Create(&ksession)
	if err != nil {
		return err
	}
	return nil
}

func DeleteByUserID(UserID int) error {
	err := sessionRepo.DeleteByUserID(UserID)
	if err != nil {
		return err
	}
	return nil
}
