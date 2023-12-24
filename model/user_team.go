package model

import "time"

type UserTeam struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	SubUserID  string    `json:"sub_user_id"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}

func (u *UserTeam) TableName() string {
	return "user_team"
}
