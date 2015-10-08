package models

import(
	"time"
)

type User struct {
		Id int `sql:"AUTO_INCREMENT"`
		Name string `sql:"size:255"`
		BookId int `sql:"index"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
}
