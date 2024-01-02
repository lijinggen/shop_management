package user_service

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/common_dto"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/repository"
	"github.com/shop_management/repository/user_repo"
	"github.com/shop_management/service"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
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

	for _, subUserRecord := range resp.List {
		if v, ok := userDetailMap[subUserRecord.SubUserId]; ok {
			subUserRecord.Name = v.Name
			subUserRecord.Phone = v.Phone
		}
	}
	return resp, nil
}

func (u *userTeamServiceImpl) AddSubUser(ctx *gin.Context, req *user_dto.AddSubUserReq) error {
	subUser, err := u.userRepo.GetByPhone(ctx, util.GetDBFromContext(ctx), req.Phone)
	if err != nil {
		return err
	}
	if subUser == nil {
		return sm_error.NewHttpError(error_code.UserNoExists)
	}
	userId, _ := ctx.Cookie("user_id")

	err = u.userTeamRepo.AddSubUser(ctx, util.GetDBFromContext(ctx), subUser.Id, userId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userTeamServiceImpl) DelSubUser(ctx *gin.Context, req *user_dto.DelSubUserReq) error {
	err := u.userTeamRepo.DelSubUser(ctx, util.GetDBFromContext(ctx), req.Id)
	if err != nil {
		return err
	}
	return nil
}
