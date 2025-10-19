package config

import (
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	database, err := gorm.Open("mysql", "root:C@stellon1209/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database")
	}
	db = database
}

func GetDB() *gorm.DB {
	return db
}
