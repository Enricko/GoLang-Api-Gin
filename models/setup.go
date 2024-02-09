package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("ricky:123qweasd@tcp(localhost:3306)/golang_api"))

	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB = database
}
