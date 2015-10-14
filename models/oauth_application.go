package models

import(
	"time"
)

type OauthApplication struct {
	Id int `sql:AUTO_INCREMENT`
	Uid string `sql:"size:255"`
	Secret string `sql:"size:255"`
	RedirectUri string `sql:"size:255"`
	Scopes string `sql:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
