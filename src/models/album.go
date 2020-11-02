package models

import (
	"gorm.io/gorm"
)

type Album struct {
	ID     uint
	Name   string
	Photos []*Photo `gorm:"many2many:album_photos;"`
	gorm.Model
}
