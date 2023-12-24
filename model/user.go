package model

import "time"

type User struct {
	BaseModel
	Id         string
	Name       string
	Email      string
	Password   string
	AvatarUrl  string
	Phone      string
	CreateTime time.Time
	ModifyTime time.Time
}

func (u *User) TableName() string {
	return "user"
}
