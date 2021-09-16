package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var dsn = "root:codelogicx101@tcp(localhost:3306)/laravel-api?charset=utf8&parseTime=True&loc=Local"
var err error

func dataMigration() {
	// Note the assignment and not initialise + assign operator
	// I made mistake here
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	DB.AutoMigrate(&Employee{})
}
