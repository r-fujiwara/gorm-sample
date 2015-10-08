package main

// if this .go execute 

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-fujiwara/gorm-sample/models"
)

func main(){
	db, err := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm-test-sample?charset=utf8&parseTime=True")

	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	db.CreateTable(&models.User{})
	//db.CreateTable(&models.Book{})
}
