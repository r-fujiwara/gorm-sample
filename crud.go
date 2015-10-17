package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-fujiwara/gorm-sample/models"
)

var db gorm.DB

func main() {
	db, err := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm_test_sample?charset=utf8mb4&parseTime=True")

	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		return
	}

	var u models.User
	// select
	db.Where("username = ?", "r-fujiwara").First(&u)
	fmt.Println(u.Username)
	// r-fujiwara


}

