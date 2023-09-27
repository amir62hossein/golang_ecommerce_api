package models

import "time"

type ROLE_ENUM string

const (
	ADMIN ROLE_ENUM = "ADMIN"
	USER  ROLE_ENUM = "USER"
)

type User struct {
	DBModel
	UserName    string    `json:"user-name" gorm:"size:255;not null"`
	Email       string    `json:"email" gorm:"unique;not null"`
	Password    string    `json:"password" gorm:"not null"`
	Phone       string    `json:"phone" gorm:"not null"`
	Role        ROLE_ENUM `json:"role" sql:"type:ENUM('USER', 'ADMIN')" gorm:"default:USER"`
	Disabled    bool      `json:"disabled" gorm:"default:false"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
