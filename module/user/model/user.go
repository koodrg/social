package usermodel

import (
	"errors"
	"social/common"
)

type User struct {
	common.SQLModel `json:"users" gorm:"users"`
	Email           string  `json:"email" gorm:"email"`
	FbId            string  `json:"fb_id" gorm:"fb_id"`
	GgId            string  `json:"gg_id" gorm:"gg_id"`
	Password        string  `json:"password" gorm:"password"`
	Salt            string  `json:"salt" gorm:"salt"`
	LastName        string  `json:"last_name" gorm:"last_name"`
	FirstName       string  `json:"first_name" gorm:"first_name"`
	Phone           string  `json:"phone" gorm:"phone"`
	Privacy         Privacy `json:"role" gorm:"privacy"`
	Avatar          string  `json:"avatar" gorm:"avatar"`
	Status          int     `json:"status" gorm:"status"`
}

func (User) TableName() string { return "users" }

type UserCreate struct {
	common.SQLModel
	Email     string `json:"email" gorm:"email"`
	FbId      string `json:"fb_id,onitempty" gorm:"fb_id"`
	GgId      string `json:"gg_id,onitempty" gorm:"gg_id"`
	LastName  string `json:"last_name" gorm:"last_name"`
	FirstName string `json:"first_name" gorm:"first_name"`
	Phone     string `json:"phone" gorm:"phone"`
	Password  string `json:"_" gorm:"password"`
	Salt      string `json:"_" gorm:"salt"`
	Role      string `json:"_" gorm:"role"`
}

func (UserCreate) TableName() string { return User{}.TableName() }

type UserUpdate struct {
	common.SQLModel
	Email     string `json:"email" gorm:"email"`
	LastName  string `json:"last_name" gorm:"last_name"`
	FirstName string `json:"first_name" gorm:"first_name"`
	Phone     string `json:"phone" gorm:"phone"`
}

func (UserUpdate) TableName() string { return User{}.TableName() }

type LoginModel struct {
	Email    string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}

type Privacy int

const (
	Private Privacy = 1
	Public  Privacy = 2
)

var (
	ErrEmailExisted = common.NewCustomError(
		errors.New("email already exists"),
		"Email already exists",
		"ErrEmailExisted",
		"BAD_REQUEST",
	)

	ErrUserHasBeenDisabled = common.NewCustomError(
		errors.New("user has been disabled"),
		"User has been disabled",
		"ErrUserHasBeenDisabled",
		"BAD_REQUEST",
	)
)
