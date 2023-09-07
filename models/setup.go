package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp:localhost:3360)/go_restapi.gin"))
	if err != nil {
		panic(err)
	}

	database.Automigrate(&Product{})

	DB = database
}
