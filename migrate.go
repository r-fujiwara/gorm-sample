package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-fujiwara/gorm-sample/models"
)

func main(){
	db, _ := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm-test-sample?charset=utf8&parseTime=True")

	db.CreateTable(&models.User{})
}
