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
	FileID      int       `json:"file_id" gorm:"type:integer"`
	Model
}

type UpdateUser struct {
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Telp     string `json:"telp"`
	Email    string `json:"email"`
	UserType string `json:"user_type"`
	FileID   int    `json:"file_id,omitempty"`
}

type ChangePassword struct {
	OldPassword     string `json:"old_password" valid:"Required"`
	NewPassword     string `json:"new_password" valid:"Required"`
	ConfirmPassword string `json:"confirm_password" valid:"Required"`
}

type AddUser struct {
	Email    string `json:"email" valid:"Required"`
	Telp     string `json:"telp"`
	Password string `json:"password"`
	Name     string `json:"name" valid:"Required"`
	IsAdmin  bool   `json:"is_admin"`
}

type LoginCapster struct {
	CapsterID      int    `json:"capster_id"`
	CapsterName    string `json:"capster_name"`
	IsActive       bool   `json:"is_active"`
	Password       string `json:"password"`
	Email          string `json:"email"`
	Telp           string `json:"telp"`
	FileID         int    `json:"file_id"`
	FileName       string `json:"file_name"`
	FilePath       string `json:"file_path"`
	BarberID       int    `json:"barber_id"`
	BarberName     string `json:"barber_name"`
	BarberIsActive bool   `json:"barber_is_active"`
	OwnerID        int    `json:"owner_id"`
	OwnerName      string `json:"owner_name"`
}

type UserList struct {
	UserID   int       `json:"user_id" gorm:"PRIMARY_KEY"`
	UserName string    `json:"user_name" gorm:"type:varchar(20)"`
	Name     string    `json:"name" gorm:"type:varchar(60);not null"`
	Telp     string    `json:"telp" gorm:"type:varchar(20)"`
	Email    string    `json:"email" gorm:"type:varchar(60)"`
	IsActive bool      `json:"is_active" gorm:"type:boolean"`
	JoinDate time.Time `json:"join_date" gorm:"type:timestamp(0)"`
	// FileID    int       `json:"file_id" gorm:"type:integer"`
	UserType string `json:"user_type" gorm:"type:varchar(10)"`
	UserEdit string `json:"user_edit" gorm:"type:varchar(20)"`
	FileOutput
}
