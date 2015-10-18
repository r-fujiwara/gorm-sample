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

	db.DB().Ping()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.OauthApplication{})
	db.CreateTable(&models.OauthAccessToken{})

	user := models.User{Username: "r-fujiwara", Password: "mino-monta"}
	// Create はdb.Debug() とかやらなくてもログに出る
	db.Create(&user)
	oauth_application := models.OauthApplication{Uid: "0123456789", Name: "normal_app", Secret: "0123456789", RedirectUri: "urn:ietf:wg:oauth:2.0:oob", Scopes: ""}
	db.Create(&oauth_application)
}

