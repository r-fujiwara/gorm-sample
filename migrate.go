package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-fujiwara/gorm-sample/models"
)

var db gorm.DB

func ResetTables(){
	db.DropTableIfExists(&models.User{})
	db.DropTableIfExists(&models.OauthApplication{})
	db.DropTableIfExists(&models.OauthAccessToken{})
}

func main() {
	db, err := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm_test_sample?charset=utf8mb4&parseTime=True")

	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
		return
	}
	// if you create table see the my qiita article
	// 何か変なエラー出て動かないのでやめた, sqlでやる
	//ResetTables()

	db.DB().Ping()

	db.CreateTable(&models.User{})
	db.CreateTable(&models.OauthApplication{})
	db.CreateTable(&models.OauthAccessToken{})

	user := models.User{Username: "r-fujiwara", Password: "mino-monta"}
	db.Create(&user)
	oauth_application := models.OauthApplication{Uid: "0123456789", Name: "native_app", Secret: "0123456789", RedirectUri: "urn:ietf:wg:oauth:2.0:oob", Scopes: ""}
	db.Create(&oauth_application)
}

