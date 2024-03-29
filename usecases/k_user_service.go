package usecases

import (
	"context"
	"errors"
	"math"
	"reflect"
	"time"

	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/interfaces/services"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/database"
	"github.com/satriaprayoga/cukurin-barber/pkg/utils"
	"github.com/satriaprayoga/cukurin-barber/token"
)

type kUserService struct {
	kuserrepo      repo.IKUserRepository
	fileRepo       repo.IFileUploadRepository
	contextTimeOut time.Duration
}

func NewKUserService(kUserRepo repo.IKUserRepository, fileRepo repo.IFileUploadRepository, cto time.Duration) services.IKUserService {
	return &kUserService{kuserrepo: kUserRepo, fileRepo: fileRepo, contextTimeOut: cto}
}

func (r *kUserService) GetByEmailSaUser(ctx context.Context, email string, usertype string) (result models.KUser, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	kuser := models.KUser{}
	result, err = r.kuserrepo.GetByAccount(email, usertype)
	if err != nil {
		return kuser, err
	}
	return result, nil

}

func (r *kUserService) ChangePassword(ctx context.Context, Claims token.Claims, DataChPwd models.ChangePassword) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	dataUser, err := r.kuserrepo.GetDataBy(Claims.UserID)
	if err != nil {
		return err
	}

	if !utils.ComparePassword(dataUser.Password, utils.GetPassword(DataChPwd.OldPassword)) {
		return errors.New("password lama anda salah")
	}

	if DataChPwd.NewPassword != DataChPwd.ConfirmPassword {
		return errors.New("password dan confirm password tidaks sama")
	}

	if utils.ComparePassword(dataUser.Password, utils.GetPassword(DataChPwd.NewPassword)) {
		return errors.New("password baru tidak boleh sama dengan yang lama")
	}

	DataChPwd.NewPassword, _ = utils.Hash(DataChPwd.NewPassword)

	err = r.kuserrepo.UpdatePasswordByEmail(dataUser.Email, DataChPwd.NewPassword)
	if err != nil {
		return err
	}
	return nil
}

func (r *kUserService) GetDataBy(ctx context.Context, Claims token.Claims, ID int) (result interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	DataOwner, err := r.kuserrepo.GetDataBy(ID)
	if err != nil {
		if err != models.ErrNotFound {
			return result, err
		}
	}

	DataFile, err := r.fileRepo.GetByID(DataOwner.FileID)
	if err != nil {
		if err != models.ErrNotFound {
			return result, err
		}
	}

	response := map[string]interface{}{
		"owner_id":   DataOwner.UserID,
		"owner_name": DataOwner.Name,
		"email":      DataOwner.Email,
		"telp":       DataOwner.Telp,
		"file_id":    DataOwner.FileID,
		"file_name":  DataFile.FileName,
		"file_path":  DataFile.FilePath,
	}
	return response, nil
}

func (r *kUserService) GetList(ctx context.Context, Claims token.Claims, queryparam models.ParamList) (result models.ResponseModelList, err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	var kUser = models.KUser{}
	if queryparam.Search != "" {
		value := reflect.ValueOf(kUser)
		types := reflect.TypeOf(&kUser)
		queryparam.Search = database.GetWhereLikeStruct(value, types, queryparam.Search, "")
	}
	result.Data, err = r.kuserrepo.GetList(queryparam)
	if err != nil {
		return result, err
	}
	result.Total, err = r.kuserrepo.Count(queryparam)
	if err != nil {
		return result, err
	}
	result.LastPage = int(math.Ceil(float64(result.Total) / float64(queryparam.PerPage)))
	result.Page = queryparam.Page

	return result, nil
}

func (r *kUserService) Create(ctx context.Context, Claims token.Claims, data *models.KUser) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	err = r.kuserrepo.Create(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *kUserService) Update(ctx context.Context, Claims token.Claims, ID int, data models.UpdateUser) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	dataUser, err := r.kuserrepo.GetByAccount(data.Email, data.UserType)
	if dataUser.UserID != ID {
		return errors.New("email sudah terdaftar")
	}
	var datas = map[string]interface{}{
		"name":    data.Name,
		"telp":    data.Telp,
		"email":   data.Email,
		"file_id": data.FileID,
	}
	datas["user_edit"] = Claims.UserID
	err = r.kuserrepo.Update(ID, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *kUserService) Delete(ctx context.Context, ID int) (err error) {
	_, cancel := context.WithTimeout(ctx, r.contextTimeOut)
	defer cancel()

	err = r.kuserrepo.Delete(ID)
	if err != nil {
		return err
	}
	return nil
}
