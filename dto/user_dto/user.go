package user_dto

import (
	"github.com/shop_management/dto/common_dto"
	"time"
)

type User struct {
	Id         string
	Name       string
	Email      string
	Password   string
	ApiKey     string
	ApiSecret  string
	AvatarUrl  string
	Phone      string
	Balance    float64
	Currency   string
	CreateTime time.Time
	ModifyTime time.Time
}

type SubUser struct {
	Id        string
	UserId    string
	SubUserId string
	Name      string
	Phone     string
}

type UserLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterUserReq struct {
	Name            string `json:"name"`
	Phone           string `json:"phone"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ModifyUserPasswordReq struct {
	UserId          string `json:"user_id"`
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type UserProfile struct {
	UserId    string
	AvatarUrl string
	Name      string
	Email     string
	Phone     string
	IsAdmin   bool
}

type GenerateApiSecret struct {
	UserId string `json:"user_id"`
	ApiKey string `json:"api_key"`
}

type GetApiSecret struct {
	ApiKey    string `json:"api_key"`
	ApiSecret string `json:"api_secret"`
}

type UserListReq struct {
	Pager   *common_dto.Pager `json:"pager"`
	Name    string            `json:"name" form:"name"`
	UserIds []string          `json:"user_ids"`
}

type UserListResp struct {
	Pager *common_dto.Pager `json:"pager"`
	List  []*User           `json:"list"`
}

type SubUserListReq struct {
	Pager  *common_dto.Pager `json:"pager"`
	UserId string            `json:"user_id"`
}

type SubUserListResp struct {
	Pager *common_dto.Pager `json:"pager"`
	List  []*SubUser        `json:"list"`
}

type AddSubUserReq struct {
	Phone string `json:"phone"`
}

type DelSubUserReq struct {
	Id string `json:"id"`
}
