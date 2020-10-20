package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func NewMySQLConnect() {
	dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DisableForeignKeyConstraintWhenMigrating: true,

	if err != nil {
		panic(err.Error())
	}

	SqlDB = gormDB
}
