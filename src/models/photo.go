package models

import (
	"gorm.io/gorm"
)

type Photo struct {
	ID          uint
	Description string
	FileName    string
	IsPublished bool
	Name        string
	Views       int64
	gorm.Model
	AuthorId uint
	Author   Author   `gorm:"foreignKey:AuthorId"`
	Alumbs   []*Album `gorm:"many2many:album_photos;"`
}
