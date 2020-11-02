package models

import (
	"gorm.io/gorm"
)

type Author struct {
	ID     uint
	Name   string
	Photos []Photo `gorm:"foreignKey:AuthorId"`
	gorm.Model
}
