package repoimpl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/logging"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"gorm.io/gorm"
)

type repoKUser struct {
	Conn *gorm.DB
}

func NewRepoKUser(Conn *gorm.DB) repo.IKUserRepository {
	return &repoKUser{Conn}
}

func (db *repoKUser) GetByAccount(account string, userType string) (result models.KUser, err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Where("email LIKE ? OR telp=? AND user_type=?", account, account, userType).Find(&result)
	logger.Query(fmt.Sprintf("%v", query))
	// logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return result, err
}

/*  */

func (db *repoKUser) GetByCapster(Account string) (result models.LoginCapster, err error) {
	var (
		logger = logging.Logger{}
	)
	query := db.Conn.Table("k_user su").Select(`su.user_id as capster_id, su."name" as capster_name,su."password",su.email ,
						su.telp ,su.file_id ,sf.file_name ,sf.file_path ,b.barber_id ,b.barber_name,su.user_input as owner_id ,so."name" as owner_name,su.is_active,
						b.is_active as barber_is_active
						`).Joins(`
						left join barber_capster bc on su.user_id = bc.capster_id`).Joins(`
						left join barber b on b.barber_id = bc.barber_id `).Joins(`
						left join file_upload sf on sf.file_id =su.file_id`).Joins(`
						left join k_user so on so.user_id::varchar = su.user_input `).Where(`
						su.user_type = 'capster' AND (su.email iLike ? OR su.telp = ?)`, Account, Account).First(&result)
	logger.Info(fmt.Sprintf("%v", query))
	err = query.Error

	if err != nil {
		//
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}

	return result, err
}

func (db *repoKUser) GetDataBy(ID int) (result *models.KUser, err error) {
	var (
		logger = logging.Logger{}
		kuser  = &models.KUser{}
	)

	query := db.Conn.Where("user_id=?", ID).Find(&kuser)
	logger.Query(fmt.Sprintf("%v", query))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return result, models.ErrNotFound
		}
		return result, err
	}
	return kuser, nil
}

func (db *repoKUser) GetList(queryparam models.ParamList) (result []*models.UserList, err error) {
	var (
		pageNum  = 0
		pageSize = settings.AppConfigSetting.App.PageSize
		sWhere   = ""
		orderBy  = queryparam.SortField
		logger   = logging.Logger{}
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

	//WHERE
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

	//end where

	query := db.Conn.Table("k_user sa ").Select(`
				sa.user_id ,sa.user_name ,
				sa."name" ,sa.telp ,
				sa.email ,sa.is_active ,
				sa.join_date ,sa.user_type ,
				sa.file_id ,sf.file_name ,
				sf.file_path ,sf.file_type
	`).Joins(`
	left join file_upload sf
	on sf.file_id = sa.file_id 
	`).Where(sWhere).Offset(pageNum).Limit(pageSize).Order(orderBy).Find(&result)
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

func (db *repoKUser) Create(data *models.KUser) error {
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

func (db *repoKUser) Update(ID int, data interface{}) error {
	var (
		logger = logging.Logger{}
		err    error
	)

	q := db.Conn.Model(&models.KUser{}).Where("user_id=?", ID).Updates(data)
	logger.Query(fmt.Sprintf("%v", q))
	err = q.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoKUser) UpdatePasswordByEmail(Email string, Password string) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Exec(`UPDATE k_user set password = ? AND email = ?`, Password, Email)
	logger.Query(fmt.Sprintf("%v", query))
	err = query.Error
	if err != nil {
		return err
	}
	return nil
}

func (db *repoKUser) Delete(ID int) error {
	var (
		logger = logging.Logger{}
		err    error
	)
	query := db.Conn.Where("user_id=?", ID).Delete(&models.KUser{})
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return err
	}
	return nil

}

func (db *repoKUser) Count(querparam models.ParamList) (result int, err error) {
	var (
		sWhere  = ""
		logger  = logging.Logger{}
		_result int64
	)

	//WHERE
	if querparam.InitSearch == "" {
		sWhere = querparam.InitSearch
	}

	if querparam.Search != "" {
		if sWhere != "" {
			sWhere += " and " + querparam.Search
		}
	}

	query := db.Conn.Model(&models.KUser{}).Where(sWhere).Count(&_result)
	logger.Query(fmt.Sprintf("%v", query)) //cath to log query string
	err = query.Error
	if err != nil {
		return 0, err
	}
	return int(_result), nil
}

func (db *repoKUser) GenUserCapster() (string, error) {
	result := ""
	row := db.Conn.Table("k_user").Select("max(user_name)").Row()
	//.Where("to_timestamp(created_on)::date = now()::date").Row()
	row.Scan(&result)
	// if result == "" {
	// 	return "error"
	// }
	// if err != nil && err != gorm.ErrRecordNotFound {
	// 	return result
	// }
	// return result, nil
	var (
		currentTime     = time.Now()
		year            = currentTime.Year()
		month       int = int(currentTime.Month())
		day             = currentTime.Day()
	)

	// result := u.repoUser.GetMaxUserCapster()
	sYear := strconv.Itoa(year)[2:]
	var sMonth string = strconv.Itoa(month)
	if len(sMonth) == 1 {
		sMonth = fmt.Sprintf("0%s", sMonth)
	}
	var sDay string = strconv.Itoa(day)
	if len(sDay) == 1 {
		sDay = fmt.Sprintf("0%s", sDay)
	}
	seqNo := "0001"
	if result == "" {
		result = fmt.Sprintf("CP%s%s%s%v", sYear, sMonth, sDay, seqNo)
	} else {
		seqNo = fmt.Sprintf("1%s", result[9:])
		iSeqno, err := strconv.Atoi(seqNo)
		if err != nil {
			return "", err
		}
		iSeqno += 1
		seqNo = strconv.Itoa(iSeqno)[1:]
		result = fmt.Sprintf("CP%s%s%s%v", sYear, sMonth, sDay, seqNo)
	}
	return result, nil
}
