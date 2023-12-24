package user_server

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/common_dto"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/po/common_po"
	"github.com/shop_management/po/user_po"
	"github.com/shop_management/service"
	"github.com/shop_management/service/user_service"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
)

type UserTeamServer struct {
	userTeamService service.UserTeamService
}

func NewUserTeamServer() *UserTeamServer {
	return &UserTeamServer{
		userTeamService: user_service.NewUserTeamServiceImpl(),
	}
}

func (u *UserTeamServer) SubUserList(ctx *gin.Context) (interface{}, error) {
	pager := &common_po.Pager{}
	err := ctx.ShouldBindQuery(pager)
	if err != nil {
		return nil, sm_error.NewHttpError(error_code.ReqParamError)
	}
	r, err := u.userTeamService.SubUserList(ctx, &user_dto.SubUserListReq{
		Pager: &common_dto.Pager{
			Page:      pager.Page,
			PageSize:  pager.PageSize,
			TotalRows: pager.TotalRows,
		},
		UserId: util.GetUserIdByCookie(ctx),
	})
	if err != nil {
		return nil, err
	}
	list := make([]*user_po.SubUser, 0)
	for _, user := range r.List {
		list = append(list, &user_po.SubUser{
			Id:    user.Id,
			Name:  user.Name,
			Phone: user.Phone,
		})
	}
	resp := &user_po.SubUserListResp{
		Pager: &common_po.Pager{
			Page:      r.Pager.Page,
			PageSize:  r.Pager.PageSize,
			TotalRows: r.Pager.TotalRows,
		},
		List: list,
	}
	return resp, nil
}
