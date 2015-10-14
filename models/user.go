package models

import(
	"time"
)

type User struct {
		Id int `sql:"AUTO_INCREMENT"`
		Username string `sql:"size:255"`
		Password string `sql:size:255`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt time.Time
}
