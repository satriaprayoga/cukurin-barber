package repoimpl

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"gorm.io/gorm"
)

type repoBarber struct {
	Conn *gorm.DB
}

func NewRepoBarber(Conn *gorm.DB) repo.IBarberRepository {
	return &repoBarber{Conn}
}

func (r *repoBarber) GetDataBy(ID int) (result *models.Barber, err error) {
	var (
		logger = logging.Logger{}
		barber = &models.Barber{}
	)
	query := r.Conn.Where("barber_id = ?", ID).Find(barber)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return barber, nil

}

func (r *repoBarber) GetDataFirst(OwnerID int, BarberID int) (result *models.Barber, err error) {
	var (
		logger = logging.Logger{}
		barber = &models.Barber{}
		sQuery = ""
	)
	if BarberID == 0 {
		sQuery = `SELECT * FROM barber where is_active = true and owner_id = ?
		order by barber_id
		limit 1`

	} else {
		sQuery = `SELECT * FROM barber where is_active = true and owner_id = ? AND barber_id = ?
		order by barber_id
		limit 1`
	}
	query := r.Conn.Raw(sQuery, OwnerID).Scan(&barber)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return barber, nil
}
func (r *repoBarber) GetList(queryparam models.ParamList) (result []*models.BarbersList, err error) {
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
			sWhere += " and lower(barber_name) LIKE ?"
		} else {
			sWhere += "lower(barber_name) LIKE ?"
		}
		query = r.Conn.Table("barber b ").Select(`
		b.barber_id,b.barber_cd,b.barber_name,
		b.address,b.latitude,b.longitude,
		b.operation_start,b.operation_end,
		b.is_active,c.file_id, c.file_name, c.file_path, c.file_type
		`).Joins(`
		left join file_upload c 
		on b.file_id=c.file_id
		`).Where(sWhere, queryparam.Search).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	} else {
		query = r.Conn.Table("barber b ").Select(`
		b.barber_id,b.barber_cd,b.barber_name,
		b.address,b.latitude,b.longitude,
		b.operation_start,b.operation_end,
		b.is_active,c.file_id, c.file_name, c.file_path, c.file_type
		`).Joins(`
		left join file_upload c 
		on b.file_id=c.file_id
		`).Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
	}
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return result, nil

}
func (r *repoBarber) Create(data *models.Barber) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := r.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repoBarber) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := r.Conn.Model(models.Barber{}).Where("barber_id = ?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repoBarber) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	// query := db.Conn.Where("barber_id = ?", ID).Delete(&models.Barber{})
	query := r.Conn.Exec("Delete From barber_collection WHERE barber_id = ?", ID)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
func (r *repoBarber) Count(queryparam models.ParamList) (result int64, err error) {
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
			sWhere += " and lower(barber_name) LIKE ?" //+ queryparam.Search
		} else {
			sWhere += "lower(barber_name) LIKE ?" //queryparam.Search
		}

		query = r.Conn.Model(&models.Barber{}).Where(sWhere, queryparam.Search).Count(&result)
	} else {
		query = r.Conn.Model(&models.Barber{}).Where(sWhere).Count(&result)
	}
	// end where

	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}

	return result, nil
}
