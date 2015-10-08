import main

import (
    "github.com/zenazn/goji"
    "github.com/zenazn/goji/web"
    "github.com/zenazn/goji/web/middleware"
    _ "github.com/go-sql-driver/mysql"
)

var db gorm.DB

func main() {
	goji.Get("/users", UserIndex)
	goji.Get("/users/:id/books", BookLists)
	goji.Post("/users/:id/books", BookLists)
	goji.Serve()
}


// controllers.users.index
func UserIndex(){
	
}

func BookLists(){
	
}

func addBook(){
	
}
