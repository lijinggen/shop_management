package user_dto

import "time"

type UserBalance struct {
	Id         string
	UserId     string
	Balance    float64
	Currency   string
	CreateTime time.Time
	ModifyTime time.Time
}

type AddUserBalanceReq struct {
	UserId          string
	Balance         float64
	BalanceCurrency string
}
