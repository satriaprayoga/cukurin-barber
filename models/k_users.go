package models

import (
	"time"
)

type KUser struct {
	UserID      int       `json:"user_id" gorm:"PRIMARY_KEY"`
	UserName    string    `json:"user_name" gorm:"type:varchar(20)"`
	Name        string    `json:"name" gorm:"type:varchar(60);not null"`
	Telp        string    `json:"telp" gorm:"type:varchar(20)"`
	Email       string    `json:"email" gorm:"type:varchar(60)"`
	IsActive    bool      `json:"is_active" gorm:"type:boolean"`
	Password    string    `json:"password" gorm:"type:varchar(150)"`
	UserType    string    `json:"user_type" gorm:"type:varchar(10)"`
	JoinDate    time.Time `json:"join_date" gorm:"type:timestamp(0)"`
	BirthOfDate time.Time `json:"birth_of_date" gorm:"type:timestamp(0)"`
	Model
}

type UpdateUser struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Telp     string `json:"telp"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password" valid:"Required"`
	NewPassword     string `json:"new_password" valid:"Required"`
	ConfirmPassword string `json:"confirm_password" valid:"Required"`
}
