package database

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func NewMySQLConnect() {
	dsn := viper.GetString("mysql.dsn")

	dsn = fmt.Sprintf(
		dsn,
		viper.Get("db.username"),
		viper.Get("db.password"),
		viper.Get("db.host"),
		viper.GetInt("db.port"),
		viper.Get("db.name"),
	)

	// dsn := "root:@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// DisableForeignKeyConstraintWhenMigrating: true,

	if err != nil {
		panic(err.Error())
	}

	SqlDB = gormDB
}
