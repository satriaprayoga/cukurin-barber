package repoimpl

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"gorm.io/gorm"
)

type repoKSession struct {
	Conn *gorm.DB
}

func NewRepoKSession(Conn *gorm.DB) repo.IKSessionRepository {
	return &repoKSession{Conn}
}

func (db *repoKSession) Create(data *models.KSession) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	q := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoKSession) GetByUserID(UserID int) (output *models.KSession, err error) {
	var (
		ksession = &models.KSession{}
		logger   = logging.Logger{}
	)
	query := db.Conn.Where("user_id=?", UserID).Find(ksession)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return ksession, nil
}

func (db *repoKSession) GetByAccount(account string) (output *models.KSession, err error) {
	var (
		ksession = &models.KSession{}
		logger   = logging.Logger{}
	)
	query := db.Conn.Where("account=?", account).Find(ksession)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return ksession, nil
}

func (db *repoKSession) GetBySessionID(SessionID string) (output *models.KSession, err error) {
	var (
		ksession = &models.KSession{}
		logger   = logging.Logger{}
	)
	query := db.Conn.Where("session_id=?", SessionID).Find(ksession)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return ksession, nil
}

func (db *repoKSession) GetByIDAndType(SessionID, sessionType string) (output *models.KSession, err error) {
	var (
		ksession = &models.KSession{}
		logger   = logging.Logger{}
	)
	query := db.Conn.Where("session_id=? AND session_type=?", SessionID, sessionType).Find(ksession)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return ksession, nil
}

func (db *repoKSession) DeleteByUserID(UserID int) (err error) {

	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Where("user_id=?", UserID).Delete(&models.KSession{})
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil

}

func (db *repoKSession) DeleteBySessionID(SessionID string) (err error) {

	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Where("session_id=?", SessionID).Delete(&models.KSession{})
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil

}
