package user_repo

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/model"
	"github.com/shop_management/repository"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"gorm.io/gorm"
)

type userTeamRepoImpl struct {
}

func NewUserTeamRepoImpl() repository.UserTeamRepo {
	return &userTeamRepoImpl{}
}

func (u *userTeamRepoImpl) List(ctx *gin.Context, db *gorm.DB, req *user_dto.SubUserListReq) (*user_dto.SubUserListResp, error) {
	var subUsers []*model.UserTeam

	offset := (req.Pager.Page - 1) * req.Pager.PageSize
	limit := req.Pager.PageSize

	if err := db.Model(&model.UserTeam{}).Where("user_id=?", req.UserId).Offset(int(offset)).Limit(int(limit)).Find(&subUsers).Error; err != nil {
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	var totalRows int64
	if err := db.Model(&model.UserTeam{}).Where("user_id=?", req.UserId).Count(&totalRows).Error; err != nil {
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	dList := make([]*user_dto.SubUser, 0)
	for _, user := range subUsers {
		dList = append(dList, &user_dto.SubUser{
			Id:        user.ID,
			SubUserId: user.SubUserID,
			UserId:    user.UserID,
		})
	}
	resp := &user_dto.SubUserListResp{
		Pager: req.Pager,
		List:  dList,
	}
	resp.Pager.TotalRows = totalRows

	return resp, nil
}

func (u *userTeamRepoImpl) AddSubUser(ctx *gin.Context, db *gorm.DB, subUserId string, userId string) error {
	err := db.Create(&model.UserTeam{
		UserID:    userId,
		SubUserID: subUserId,
	}).Error
	if err != nil {
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userTeamRepoImpl) DelSubUser(ctx *gin.Context, db *gorm.DB, id string) error {
	err := db.Where("id=?", id).Delete(&model.UserTeam{}).Error
	if err != nil {
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}
