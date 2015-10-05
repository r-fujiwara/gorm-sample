package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/r-fujiwara/gorm-sample/models"
)

func main(){
	db, _ := gorm.Open("mysql", "root:@/gorm?charset=utf8&parseTime=True")
	db.CreateTable(&models.User{})
}
