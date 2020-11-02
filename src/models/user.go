package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Name      string
	Email     *string
	Age       uint8
	Username  string
	Password  string `json:"-"`
	FirstName string
	LastName  string
	gorm.Model
}
