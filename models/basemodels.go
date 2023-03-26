package models

import (
	"errors"
	"time"
)

// Model :
type Model struct {
	UserInput string    `json:"user_input" gorm:"type:varchar(20)"`
	UserEdit  string    `json:"user_edit" gorm:"type:varchar(20)"`
	TimeInput time.Time `json:"time_input" gorm:"type:timestamp(0) without time zone;default:now()"`
	TimeEdit  time.Time `json:"time_edit" gorm:"type:timestamp(0) without time zone;default:now()"`
	// IDCreated  int `json:"id_created"`
}

type ParamList struct {
	Page       int    `json:"page" valid:"Required"`
	PerPage    int    `json:"per_page" valid:"Required"`
	Search     string `json:"search,omitempty"`
	InitSearch string `json:"init_search,omitempty"`
	SortField  string `json:"sort_field,omitempty"`
}

type ParamListGeo struct {
	GeoBarber
	ParamList
}

// ResponseModelList :
type ResponseModelList struct {
	Page         int         `json:"page"`
	Total        int         `json:"total"`
	LastPage     int         `json:"last_page"`
	DefineSize   string      `json:"define_size"`
	DefineColumn string      `json:"define_column"`
	AllColumn    string      `json:"all_column"`
	Data         interface{} `json:"data"`
	Msg          string      `json:"message"`
}

var (
	// ErrInternalServerError : will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound : will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict : will throw if the current action already exists
	ErrConflict = errors.New("Your Item already exist")
	// ErrBadParamInput : will throw if the given request-body or params is not valid
	ErrBadParamInput = errors.New("Given Param is not valid")

	Unauthorized = errors.New("UnAuthorized")

	ErrInvalidLogin = errors.New("Invalid User Or Password")
)
