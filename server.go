package main

import (
		"fmt"
		"log"
		"net/http"
    "github.com/zenazn/goji"
		"io/ioutil"
		"encoding/json"
		//"encoding/xml"
    //"github.com/zenazn/goji/web"
		"github.com/jinzhu/gorm"
    _ "github.com/go-sql-driver/mysql"
)

type Params struct {
	Client_Id string
	Client_Secret string
	Grant_Type string
	UserName string
	Password string
}

var db gorm.DB

func main() {
	db, err := gorm.Open("mysql", "test-gorm-user:test-gorm-user@/gorm_test_sample?charset=utf8mb4&parseTime=True")

	db.LogMode(true)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	goji.Post("/oauth/token", TokenCreate)
	goji.Serve()
}

func TokenCreate(w http.ResponseWriter, r *http.Request){
	// read from json body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("************")
	}
	// convert to string
	raw := string(body)
	log.Println("*********raw string**************")
	log.Println(raw)

	// convert string to json
	var params Params
	err = json.Unmarshal([]byte(raw), &params)
	if err == nil {
		log.Println("*********converted json**************")
		fmt.Printf("%+v\n", params)
		// fetch json value by key
		log.Println("*********client_id**************")
		fmt.Printf("%+v\n", params.Client_Id)
	} else {
		fmt.Println(err)
		fmt.Printf("%+v\n", params)
	}
	
}
