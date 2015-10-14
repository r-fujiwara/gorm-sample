package models

import(
	"time"
)

type OauthAccessToken struct {
	Id int `sql:AUTO_INCREMENT`
	ResourceOwnerId int
	ApplicationId int
	Token string `sql:"size:255"`
	RefreshToken string `sql:"size:255"`
	RevokeAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
