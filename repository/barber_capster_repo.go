package repoimpl

import (
	"fmt"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"gorm.io/gorm"
)

type repoBarberCapster struct {
	Conn *gorm.DB
}

func NewRepoBarberCapster(Conn *gorm.DB) repo.IBarberCapsterRepository {
	return &repoBarberCapster{Conn}
}

func (db *repoBarberCapster) GetDataBy(ID int) (result *models.BarberCapster, err error) {
	var (
		logger = logging.Logger{}
		bc     = &models.BarberCapster{}
	)
	query := db.Conn.Where("barber_id=?", ID).Find(bc)
	logger.Query(fmt.Sprintf("%v", query))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, models.ErrNotFound
		}
		return nil, err
	}
	return bc, nil
}

func (db *repoBarberCapster) GetList(queryparam models.ParamList) (result []*models.CapsterList, err error) {
	var (
		pageNum  = 0
		pageSize = settings.AppConfigSetting.App.PageSize
		sWhere   = ""
		logger   = logging.Logger{}
		orderBy  = queryparam.SortField
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
			sWhere += " and " + queryparam.Search
		} else {
			sWhere += queryparam.Search
		}
	}

	// end where
	query := db.Conn.Table("k_user").Select("k_user.user_id as capster_id,k_user.name,k_user.is_active,file_upload.file_id,file_upload.file_id,file_upload.file_name,file_upload.file_path,file_upload.file_type, '0' as rating").Joins("left join file_upload ON file_upload.file_id = k_user.file_id").Joins("inner join barber_capster ON barber_capster.capster_id=k_user.user_id").Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
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

func (db *repoBarberCapster) Create(data *models.BarberCapster) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Create(data)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoBarberCapster) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Model(models.BarberCapster{}).Where("barber_id=?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoBarberCapster) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Exec("Delete FROM barber_capster WHERE barber_id=?", ID)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoBarberCapster) DeleteByCapster(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Exec("Delete FROM barber_capster WHERE capster_id=?", ID)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}
