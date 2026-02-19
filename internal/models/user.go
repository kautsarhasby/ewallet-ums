package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"column:username;type:varchar(20);unique;" validate:"required"`
	Email       string    `json:"email" gorm:"column:email;type:varchar(100);unique;" validate:"required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;type:varchar(15);" validate:"required"`
	Fullname    string    `json:"fullname" gorm:"column:fullname;type:varchar(100);" validate:"required"`
	Address     string    `json:"address" gorm:"column:address;type:text"`
	Dob         string    `json:"dob" gorm:"column:dob;type:date"`
	Password    string    `json:"password" gorm:"column:password;type:varchar(255);" validate:"required,min=6"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (*User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  uint      `gorm:"primarykey"`
	UserID              int       `json:"user_id" gorm:"type:int;" validate:"required"`
	Token               string    `json:"token" gorm:"type:varchar(255);" validate:"required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:varchar(255);" validate:"required"`
	TokenExpired        time.Time `json:"token_expired" validate:"required"`
	RefreshTokenExpired time.Time `json:"refresh_token_expired" validate:"required"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (*UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
