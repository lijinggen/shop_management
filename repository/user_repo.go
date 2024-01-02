package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/user_dto"
	"gorm.io/gorm"
)

type UserRepo interface {
	Add(ctx *gin.Context, db *gorm.DB, dto *user_dto.User) error
	GetById(ctx *gin.Context, db *gorm.DB, id string) (*user_dto.User, error)
	GetByEmail(ctx *gin.Context, db *gorm.DB, email string) (*user_dto.User, error)
	GetByPhone(ctx *gin.Context, db *gorm.DB, phone string) (*user_dto.User, error)
	ModifyPassword(ctx *gin.Context, db *gorm.DB, id string, password string) error
	SaveProfile(ctx *gin.Context, db *gorm.DB, req *user_dto.UserProfile) error
	Delete(ctx *gin.Context, db *gorm.DB, userId string) error
	SaveSecret(ctx *gin.Context, db *gorm.DB, userId, key, secret string) error
	List(ctx *gin.Context, db *gorm.DB, req *user_dto.UserListReq) ([]*user_dto.User, error)
	Save(ctx *gin.Context, db *gorm.DB, req *user_dto.User) error
	GetAll(ctx *gin.Context, db *gorm.DB) ([]*user_dto.User, error)
	GetBySecret(ctx *gin.Context, db *gorm.DB, secret string) (*user_dto.User, error)
}

type UserTeamRepo interface {
	List(ctx *gin.Context, db *gorm.DB, req *user_dto.SubUserListReq) (*user_dto.SubUserListResp, error)
	AddSubUser(ctx *gin.Context, db *gorm.DB, subUserId string, userId string) error
	DelSubUser(ctx *gin.Context, db *gorm.DB, id string) error
}
