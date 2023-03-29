package repo

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"gorm.io/gorm"
)

type repoPaket struct {
	Conn *gorm.DB
}

func NewRepoPaket(Conn *gorm.DB) IPaketRepository {
	return &repoPaket{Conn}
}

func (db *repoPaket) GetDataBy(ID int) (result *models.Paket, err error) {
	var (
		logger = logging.Logger{}
		mPaket = &models.Paket{}
	)
	query := db.Conn.Where("paket_id = ? ", ID).Find(mPaket)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return mPaket, nil
}

func (db *repoPaket) GetList(queryparam models.ParamList) (result []*models.Paket, err error) {

	var (
		pageNum  = 0
		pageSize = settings.AppConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = queryparam.SortField
		query    *gorm.DB
	)
	// pagination
	if queryparam.Page > 0 {
		pageNum = (queryparam.Page - 1) * queryparam.PerPage
	}
	if queryparam.PerPage > 0 {
		pageSize = queryparam.PerPage
	}
	//end pagination

	// Order
	if queryparam.SortField != "" {
		orderBy = queryparam.SortField
	}
	//end Order by

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and (lower(paket_name) LIKE ? OR lower(descs) LIKE ?)" //+ queryparam.Search
		} else {
			sWhere += "(lower(barber_name) LIKE ? OR lower(descs) LIKE ?)" //queryparam.Search
		}
		query = db.Conn.Where(sWhere, queryparam.Search, queryparam.Search).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	} else {
		query = db.Conn.Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}

	// end where

	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}
	return result, nil
}
func (db *repoPaket) Create(data *models.Paket) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoPaket) Update(ID int, data map[string]interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.Paket{}).Where("paket_id = ?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoPaket) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Where("paket_id = ?", ID).Delete(&models.Paket{})
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (db *repoPaket) Count(queryparam models.ParamList) (result int64, err error) {
	var (
		sWhere = ""
		logger = logging.Logger{}
		query  *gorm.DB
	)
	result = 0

	// WHERE
	if queryparam.InitSearch != "" {
		sWhere = queryparam.InitSearch
	}

	if queryparam.Search != "" {
		if sWhere != "" {
			sWhere += " and (lower(paket_name) LIKE ? OR lower(descs) LIKE ?)" //+ queryparam.Search
		} else {
			sWhere += "(lower(barber_name) LIKE ? OR lower(descs) LIKE ?)" //queryparam.Search
		}
		query = db.Conn.Model(&models.Paket{}).Where(sWhere, queryparam.Search, queryparam.Search).Count(&result)
	} else {
		query = db.Conn.Model(&models.Paket{}).Where(sWhere).Count(&result)
	}
	// end where

	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
