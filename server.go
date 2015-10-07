import main

import (
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
    _ "github.com/go-sql-driver/mysql"
)

var db gorm.DB

func main() {
	goji.Serve()
}

func init(){
	db, _ = gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm-test-sample?charset=utf8&parseTime=True")
}

