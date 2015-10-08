package models

import(
	"time"
)

type User struct {
		id int64 `gorm:"primary_key"`
		name string `sql:"size:255"`
		book_id int64 `sql:"index"`
		created_at time.Time
		updated_at time.Time
		deleted_at time.Time
}
