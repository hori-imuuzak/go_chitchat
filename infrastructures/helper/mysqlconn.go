package helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBOpen() gorm.DB, error {
	return gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}

func DBClose(db *gorm.DB) {
	db.Close()
}