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

func (F *Forgot) SendForgot() error {
	subjectEmail := "Permintaan Lupa Password"
	fmt.Printf("%s", subjectEmail)
	err := email.SendEmail(F.Email, subjectEmail, getInfoLoginBodyForgot(F))
	if err != nil {
		return err
	}
	return nil
}

func getInfoLoginBodyForgot(F *Forgot) string {
	verifyHTML := email.VerifyCode

	verifyHTML = strings.ReplaceAll(verifyHTML, `{Name}`, F.Name)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{Email}`, F.Email)
	verifyHTML = strings.ReplaceAll(verifyHTML, `{OTP}`, F.OTP)
	return verifyHTML
}
