package models

type LoginForm struct {
	Account  string `json:"account" valid:"Required"`
	Password string `json:"pwd" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
	//FcmToken string `json:"fcm_token" valid:"Required"`
}

// RegisterForm :
type RegisterForm struct {
	EmailAddr string `json:"email,omitempty" valid:"Email;Required"`
}

type VerifyForm struct {
	Account    string `json:"account" valid:"Required"`
	VerifyCode string `json:"verify_code" valid:"Required"`
}

// ResetPasswd :
type ResetPasswd struct {
	Account       string `json:"account" valid:"Required"`
	Passwd        string `json:"pwd" valid:"Required"`
	UserType      string `jsong:"user_type" valid:"Required"`
	ConfirmPasswd string `json:"confirm_pwd" valid:"Required"`
}

// ForgotForm :
type ForgotForm struct {
	Account  string `json:"account" valid:"Required"`
	UserType string `json:"user_type" valid:"Required"`
}
