package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// The point of this file will be to return a variable called db so other
// files can interact with the persistence layer
var (
	db *gorm.DB
)

func Connect() {

	d, err := gorm.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}


