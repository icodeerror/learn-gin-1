package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("mysql", "root:@(localhost)/gin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to the database!")
	}
	database.AutoMigrate(&Book{})

	DB = database
}
