package user_server

import (
	"github.com/gin-gonic/gin"
	"github.com/shop_management/po/common_po"
	"github.com/shop_management/po/user_po"
	"github.com/shop_management/server/assembly/user_assembly"
	"github.com/shop_management/service"
	"github.com/shop_management/service/user_service"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
)

type UserServer struct {
	userService service.UserService
}

func NewUserServer() *UserServer {
	return &UserServer{
		userService: user_service.NewUserServiceImpl(),
	}
}

func (u *UserServer) Login(ctx *gin.Context) (interface{}, error) {
	loginReq := &user_po.UserLogin{}
	err := ctx.ShouldBindJSON(loginReq)
	if err != nil {
		return nil, sm_error.NewHttpError(error_code.ReqParamError)
	}

	id, err := u.userService.Login(ctx, user_assembly.ConvertULPoToDto(loginReq))
	return &user_po.UserLoginResponse{UserId: id}, err
}

func (u *UserServer) Register(ctx *gin.Context) (interface{}, error) {
	registerReq := &user_po.RegisterUserReq{}
	err := ctx.ShouldBindJSON(registerReq)
	if err != nil {
		return nil, sm_error.NewParamHttpError(err)
	}
	err = u.userService.Register(ctx, user_assembly.ConvertRURPoToDto(registerReq))
	return &common_po.CommonResp{}, err
}

func (u *UserServer) GetUserProfile(ctx *gin.Context) (interface{}, error) {
	profile, err := u.userService.GetUserProfile(ctx, ctx.Query("user_id"))
	return user_assembly.ConvertUPDtoToPo(profile), err
}
