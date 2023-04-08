package usecases

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/satriaprayoga/cukurin-barber/interfaces/repo"
	"github.com/satriaprayoga/cukurin-barber/interfaces/services"
	"github.com/satriaprayoga/cukurin-barber/models"
	"github.com/satriaprayoga/cukurin-barber/pkg/sessions"
	"github.com/satriaprayoga/cukurin-barber/pkg/settings"
	"github.com/satriaprayoga/cukurin-barber/pkg/utils"
	"github.com/satriaprayoga/cukurin-barber/token"
)

type authService struct {
	repoKUser repo.IKUserRepository
	//repoKSession   repo.IKSessionRepository
	repoFileUpload repo.IFileUploadRepository
	contextTimeOut time.Duration
}

func NewAuthService(a repo.IKUserRepository, f repo.IFileUploadRepository /*b repo.IKSessionRepository,*/, timeout time.Duration) services.IAuthService {
	return &authService{repoKUser: a /*repoKSession: b,*/, repoFileUpload: f, contextTimeOut: timeout}
}

func (a *authService) Logout(ctx context.Context, claims token.Claims) error {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	err := sessions.DeleteByUserID(claims.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (a *authService) Login(ctx context.Context, dataLogin *models.LoginForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		DataOwner        = models.KUser{}
		DataCapster      = models.LoginCapster{}
		isBarber    bool = true
		response         = map[string]interface{}{}
		expireToken      = settings.AppConfigSetting.JWTExpired
		canChange   bool = false
		//ksession    models.KSession
	)

	if dataLogin.UserType == "owner" {
		DataOwner, err = a.repoKUser.GetByAccount(dataLogin.Account, dataLogin.UserType)
		if DataOwner.UserType == "" && err == models.ErrNotFound {
			return nil, errors.New("email anda belum terdaftar")
		} else {
			if DataOwner.UserType == "capster" {
				DataCapster, err = a.repoKUser.GetByCapster(dataLogin.Account)
				if err != nil {
					return nil, errors.New("email anda belum terdaftar")
				}
				if !DataCapster.IsActive {
					return nil, errors.New("account anda belum aktif")
				}
				if DataCapster.BarberID == 0 {
					return nil, errors.New("anda belum terdaftar di barber manapun")
				}
				if !DataCapster.BarberIsActive {
					return nil, errors.New("saat ini barber anda tidak aktif")
				}
				isBarber = false
			} else {
				DataCapster, err = a.repoKUser.GetByCapster(dataLogin.Account)
				if DataCapster.Email != "" && DataCapster.Email == dataLogin.Account {
					canChange = true
				}
			}
			if !DataOwner.IsActive {
				return nil, errors.New("account anda belum aktif")
			}
		}
	} else {
		isBarber = false
		DataCapster, err = a.repoKUser.GetByCapster(dataLogin.Account)
		if err != nil {
			return nil, errors.New("email anda belum terdaftar")
		}
		if !DataCapster.IsActive {
			return nil, errors.New("account anda belum aktif. Silahkan hubungi pemilik Barber")
		}

		if DataCapster.BarberID == 0 {
			return nil, errors.New("anda belum terhubung dengan Barber, Silahkan hubungi pemilik Barber")
		}

		if !DataCapster.BarberIsActive {
			return nil, errors.New("saat ini barber anda sedang tidak aktif")
		}

		DataOwner, err = a.repoKUser.GetByAccount(dataLogin.Account, dataLogin.UserType)
		if DataOwner.Email != "" && DataOwner.Email == dataLogin.Account {
			canChange = true
		}
	}

	if isBarber {
		if !utils.ComparePassword(DataOwner.Password, utils.GetPassword(dataLogin.Password)) {
			return nil, errors.New("password yang anda masukkan salah")
		}
		DataFile, err := a.repoFileUpload.GetByID(DataOwner.FileID)
		if err != nil {
			return nil, errors.New("foto tidak ditemukan")
		}
		sessionID := uuid.New().String()
		jwtToken, err := token.GenerateToken(sessionID, DataOwner.UserID, DataOwner.UserName, DataOwner.UserType)
		if err != nil {
			return nil, err
		}
		err = sessions.CreateSession(sessionID, "auth", dataLogin.Account, DataOwner.UserID, time.Now().Add(time.Hour*time.Duration(expireToken)))
		if err != nil {
			return nil, err
		}
		restUser := map[string]interface{}{
			"owner_id":   DataOwner.UserID,
			"owner_name": DataOwner.Name,
			"email":      DataOwner.Email,
			"telp":       DataOwner.Telp,
			"file_id":    DataOwner.FileID,
			"file_name":  DataFile.FileName,
			"file_path":  DataFile.FilePath,
		}
		response = map[string]interface{}{
			"token":      jwtToken,
			"data_owner": restUser,
			"user_type":  "barber",
			"can_change": canChange,
		}

	} else {
		if !utils.ComparePassword(DataCapster.Password, utils.GetPassword(dataLogin.Password)) {
			return nil, errors.New("password yang anda masukkan salah")
		}
		sessionID := uuid.New().String()
		jwtToken, err := token.GenerateToken(sessionID, DataCapster.CapsterID, DataCapster.CapsterName, "capster")
		if err != nil {
			return nil, err
		}
		err = sessions.CreateSession(sessionID, "auth", dataLogin.Account, DataCapster.CapsterID, time.Now().Add(time.Hour*time.Duration(expireToken)))
		if err != nil {
			return nil, err
		}

		restUser := map[string]interface{}{
			"owner_id":     DataCapster.OwnerID,
			"owner_name":   DataCapster.OwnerName,
			"barber_id":    DataCapster.BarberID,
			"barber_name":  DataCapster.BarberName,
			"capster_id":   DataCapster.CapsterID,
			"email":        DataCapster.Email,
			"telp":         DataCapster.Telp,
			"capster_name": DataCapster.CapsterName,
			"file_id":      DataCapster.FileID,
			"file_name":    DataCapster.FileName,
			"file_path":    DataCapster.FilePath,
		}
		response = map[string]interface{}{
			"token":        jwtToken,
			"data_capster": restUser,
			"user_type":    "capster",
			"can_change":   canChange,
		}
	}

	return response, nil
}

func (a *authService) ForgotPassword(ctx context.Context, dataForgt *models.ForgotForm) (result string, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	DataUser, err := a.repoKUser.GetByAccount(dataForgt.Account, dataForgt.UserType)
	if err != nil {
		return "", errors.New("akun tidak valid")
	}

	if DataUser.UserName == "" {
		return "", errors.New("akun tidak valid")
	}

	GenOTP := utils.GenerateNumber(4)

	mailService := &Forgot{
		Email: DataUser.Email,
		Name:  DataUser.Name,
		OTP:   GenOTP,
	}

	data, err := sessions.GetSession(GenOTP)
	if err == nil {
		fmt.Printf("deleting session :%v", data)
		sessions.DeleteBySessionID(GenOTP)
	}
	err = sessions.CreateSession(GenOTP, "forgot", DataUser.Email, DataUser.UserID, time.Now().Add(time.Hour*time.Duration(24)))
	if err != nil {
		return "", err
	}

	go mailService.SendForgot()

	//DataUser.Password, _ = utils.Hash(GenOTP)
	//err = a.repoKUser.UpdatePasswordByEmail(dataForgt.Account, DataUser.Password)
	//if err != nil {
	//	return "", err
	//}

	//mailservice

	return GenOTP, nil
}

func (a *authService) ResetPassword(ctx context.Context, dataReset *models.ResetPasswd) (err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	if dataReset.Passwd != dataReset.ConfirmPasswd {
		return errors.New("password dan confirm Password harus sama")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataReset.Account, dataReset.UserType)
	if err != nil {
		return err
	}

	if utils.ComparePassword(DataUser.Password, utils.GetPassword(dataReset.Passwd)) {
		return errors.New("password baru tidak boleh sama dengan yang lama")
	}

	DataUser.Password, _ = utils.Hash(dataReset.Passwd)

	err = a.repoKUser.UpdatePasswordByEmail(dataReset.Account, DataUser.Password)
	if err != nil {
		return err
	}

	return nil

}

func (a *authService) Register(ctx context.Context, dataRegister models.RegisterForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		User models.KUser
		//ksession models.KSession
	)

	CekData, err := a.repoKUser.GetByAccount(dataRegister.EmailAddr, "owner")
	if CekData.Email == dataRegister.EmailAddr {
		return output, errors.New("email sudah terdaftar")
	}

	GenPassword := utils.GenerateCode(4)
	User.Name = ""

	User.UserType = "owner"
	User.UserEdit = "cukur_in"
	User.UserInput = "cukur_in"
	User.Email = dataRegister.EmailAddr
	User.IsActive = true
	User.JoinDate = time.Now()

	User.Password, _ = utils.Hash(GenPassword)

	err = a.repoKUser.Create(&User)
	if err != nil {
		return output, err
	}
	mUser := map[string]interface{}{
		"user_input": strconv.Itoa(User.UserID),
		"user_edit":  strconv.Itoa(User.UserID),
	}
	err = a.repoKUser.Update(User.UserID, mUser)
	if err != nil {
		return output, err
	}

	mailService := &Register{
		Email:      User.Email,
		Name:       User.Email,
		PasswordCd: GenPassword,
	}

	go mailService.SendRegister()

	// if CekData.UserID > 0 && !CekData.IsActive {
	// 	CekData.Name = User.Name
	// 	CekData.Password = User.Password
	// 	CekData.JoinDate = User.JoinDate
	// 	CekData.UserType = User.UserType
	// 	CekData.IsActive = User.IsActive
	// 	CekData.Email = User.Email
	// 	err = a.repoKUser.Update(CekData.UserID, CekData)
	// 	if err != nil {
	// 		return output, err
	// 	}
	// } else {
	// 	User.UserEdit = dataRegister.UserName
	// 	User.UserInput = dataRegister.UserName
	// 	err = a.repoKUser.Create(&User)
	// 	if err != nil {
	// 		return output, err
	// 	}
	// }

	// GenCode := utils.GenerateNumber(4)
	//ksession.SessionID = GenCode
	//ksession.UserID = User.UserID
	//ksession.SessionType = "register"
	//ksession.ExpiresAt = time.Now().Add(time.Hour * time.Duration(24))
	//ksession.Account = User.Email
	err = sessions.CreateSession(GenPassword, "register", User.Email, User.UserID, time.Now().Add(time.Hour*time.Duration(24)))
	//err = a.repoKSession.Create(&ksession)
	if err != nil {
		a.repoKUser.Delete(User.UserID)
		return nil, err
	}
	out := map[string]interface{}{
		"gen_password": GenPassword,
	}

	return out, nil
}

func (a *authService) VerifyRegisterLogin(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	var (
		expireToken = settings.AppConfigSetting.JWTExpired
		//ksession    models.KSession
	)

	data, err := sessions.GetSessionByAccount(dataVerify.Account)
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}
	if data.SessionID != dataVerify.VerifyCode {
		return nil, errors.New("otp yang anda masukkan salah")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataVerify.Account, "user")
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}

	sessions.DeleteByUserID(DataUser.UserID)

	mUser := map[string]interface{}{
		"is_active": true,
	}

	err = a.repoKUser.Update(DataUser.UserID, mUser)
	if err != nil {
		return output, err
	}

	sessionID := uuid.New().String()

	jwtToken, err := token.GenerateToken(sessionID, DataUser.UserID, DataUser.UserName, DataUser.UserType)
	if err != nil {
		return nil, err
	}
	//ksession.SessionID = sessionID
	//ksession.UserID = DataUser.UserID
	//ksession.Account = DataUser.Email
	//ksession.ExpiresAt = time.Now().Add(time.Hour * time.Duration(expireToken))
	//ksession.SessionType = "auth"

	err = sessions.CreateSession(sessionID, "auth", DataUser.Email, DataUser.UserID, time.Now().Add(time.Hour*time.Duration(expireToken)))

	//err = a.repoKSession.Create(&ksession)
	if err != nil {
		return nil, err
	}
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.UserName,
		"user_type": DataUser.UserType,
		"name":      DataUser.Name,
	}
	response := map[string]interface{}{
		"token":     jwtToken,
		"data_user": restUser,
	}

	return response, nil
}

func (a *authService) VerifyRegister(ctx context.Context, dataVerify *models.VerifyForm) (output interface{}, err error) {
	_, cancel := context.WithTimeout(ctx, a.contextTimeOut)
	defer cancel()

	data, err := sessions.GetSessionByAccount(dataVerify.Account)
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}
	if data.SessionID != dataVerify.VerifyCode {
		return nil, errors.New("otp yang anda masukkan salah")
	}

	DataUser, err := a.repoKUser.GetByAccount(dataVerify.Account, "user")
	if err != nil {
		return nil, errors.New("akun yang anda masukkan salah")
	}

	sessions.DeleteByUserID(DataUser.UserID)

	mUser := map[string]interface{}{
		"is_active": true,
	}

	err = a.repoKUser.Update(DataUser.UserID, mUser)
	if err != nil {
		return output, err
	}
	restUser := map[string]interface{}{
		"user_id":   DataUser.UserID,
		"email":     DataUser.Email,
		"telp":      DataUser.Telp,
		"user_name": DataUser.UserName,
		"user_type": DataUser.UserType,
		"name":      DataUser.Name,
	}

	return restUser, nil
}
