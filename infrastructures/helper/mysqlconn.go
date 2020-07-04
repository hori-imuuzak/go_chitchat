package helper

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func DBOpen() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "user:password@/master?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}

	db.LogMode(true)

	return db, nil
}

func DBClose(db *gorm.DB) {
	db.Close()
}
