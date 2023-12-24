package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/common_dto"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/repository"
	"github.com/shop_management/repository/user_repo"
	"github.com/shop_management/service"
	"github.com/shop_management/util"
)

type userTeamServiceImpl struct {
	userTeamRepo repository.UserTeamRepo
	userRepo     repository.UserRepo
}

func NewUserTeamServiceImpl() service.UserTeamService {
	return &userTeamServiceImpl{
		userTeamRepo: user_repo.NewUserTeamRepoImpl(),
		userRepo:     user_repo.NewUserRepoImpl(),
	}
}

func (u *userTeamServiceImpl) SubUserList(ctx *gin.Context, req *user_dto.SubUserListReq) (*user_dto.SubUserListResp, error) {
	resp, err := u.userTeamRepo.List(ctx, util.GetDBFromContext(ctx), req)
	if err != nil {
		return nil, err
	}
	userIds := make([]string, 0)
	for _, user := range resp.List {
		userIds = append(userIds, user.Id)
	}
	userDetailList, err := u.userRepo.List(ctx, util.GetDBFromContext(ctx), &user_dto.UserListReq{
		Pager: &common_dto.Pager{
			Page:     1,
			PageSize: req.Pager.PageSize,
		},
		UserIds: userIds,
	})
	if err != nil {
		return nil, err
	}
	userDetailMap := make(map[string]*user_dto.User)
	for _, user := range userDetailList {
		userDetailMap[user.Id] = user
	}

	for _, user := range resp.List {
		if v, ok := userDetailMap[user.Id]; ok {
			user.Name = v.Name
			user.Phone = v.Phone
		}
	}
	return resp, nil
}
