package service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/user_dto"
)

type UserService interface {
	Register(ctx *gin.Context, req *user_dto.RegisterUserReq) error
	Login(ctx *gin.Context, req *user_dto.UserLogin) (string, error)
	//ModifyPwd(ctx *gin.Context, req *user_dto.ModifyUserPasswordReq) error
	GetUserProfile(ctx *gin.Context, userId string) (*user_dto.UserProfile, error)
	//SaveUserProfile(ctx *gin.Context, req *user_dto.UserProfile) error
	//Deactivate(ctx *gin.Context, userId string) error
	//GenerateApiSecret(ctx *gin.Context, req *user_dto.GenerateApiSecret) error
	//GetApiSecret(ctx *gin.Context, userId string) (*user_dto.GetApiSecret, error)
	//UserList(ctx *gin.Context, req *user_dto.UserListReq) (*user_dto.UserListResp, error)
	//SaveUser(ctx *gin.Context, req *user_dto.User) error
	//GetUserIdBySecret(ctx *gin.Context, key string, secret string) (string, error)
}

type UserTeamService interface {
	SubUserList(ctx *gin.Context, req *user_dto.SubUserListReq) (*user_dto.SubUserListResp, error)
}
