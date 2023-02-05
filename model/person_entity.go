package model

import (
	"gorm.io/gorm"
)

type Persons struct {
	gorm.Model
	UserName string `gorm:"column:user_name;type:varchar(255);not null;unique"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
	Name     string `gorm:"column:name;type:varchar(255);not null"`
	Email    string `gorm:"column:email;type:varchar(255);not null"`
	Phone    int    `gorm:"column:phone;type:integer;not null"`
}
