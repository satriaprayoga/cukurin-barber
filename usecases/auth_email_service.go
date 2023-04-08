package usecases

import (
	"fmt"
	"strings"

	"github.com/satriaprayoga/cukurin-barber/pkg/email"
)

type Forgot struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	OTP   string `json:"otp"`
}

type Register struct {
	Email      string `json:"email"`
	Name       string `json:"name"`
	PasswordCd string `json:"generate_no"`
}

func (F *Forgot) SendForgot() error {
	subjectEmail := "Permintaan Lupa Password"
	fmt.Printf("%s", subjectEmail)
	verifyHTML := email.VerifyCode

	verifyHTML = strings.ReplaceAll(verifyHTML, `{Name}`, F.Name)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{Email}`, F.Email)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{OTP}`, F.OTP)

	err := email.SendEmail(F.Email, subjectEmail, verifyHTML)
	if err != nil {
		return err
	}
	return nil
}

func (R *Register) SendRegister() error {
	subjectEmail := "Informasi Login"
	fmt.Printf(subjectEmail)
	verifyHTML := email.SendRegister

	verifyHTML = strings.ReplaceAll(verifyHTML, `{Name}`, R.Name)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{Email}`, R.Email)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{PasswordCode}`, R.PasswordCd)
	err := email.SendEmail(R.Email, subjectEmail, verifyHTML)
	if err != nil {
		return err
	}
	return nil
}
