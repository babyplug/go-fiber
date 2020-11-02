package models

import (
	"gorm.io/gorm"
)

type PhotoMetadata struct {
	ID          uint
	Height      int64
	Width       int64
	Orientation string `gorm:"size:255"`
	Compressed  int64  `gorm:"type:tinyInt"`
	Comment     string `gorm:"size:500"`
	PhotoId     uint
	Photo       Photo `gorm:"foreignKey:PhotoId"`
	gorm.Model
}
