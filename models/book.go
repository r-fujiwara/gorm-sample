package models

import(
	"time"
)

type Book struct {
    id int64 `gorm:"primary_key"`
    name string `sql:"size:255"`
		author string `sql:"size:255"`
    created_at time.Time
    updated_at time.Time
    deleted_at time.Time
}

