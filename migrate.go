package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/r-fujiwara/gorm-sample/models"
)

var db gorm.DB

func main(){
	// create databaseしてくれるわけではない
	// 大人しくCreate Databaseするのが無難か？
	db, err := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm_test_sample?charset=utf8mb4&parseTime=True")

	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		return
	}
	
	db.DB().Ping()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.OauthApplication{})
	db.CreateTable(&models.OauthAccessToken{})
}

func ResetTables(){
	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.OauthApplication{})
	db.DropTableIfExists(&models.OauthAccessToken{})
}
