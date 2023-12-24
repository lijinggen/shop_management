package user_assembly

import (
	"github.com/jinzhu/copier"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/po/user_po"
)

func ConvertRURPoToDto(req *user_po.RegisterUserReq) *user_dto.RegisterUserReq {
	convertRes := &user_dto.RegisterUserReq{}
	_ = copier.Copy(convertRes, req)
	return convertRes
}

func ConvertULPoToDto(req *user_po.UserLogin) *user_dto.UserLogin {
	convertRest := &user_dto.UserLogin{}
	_ = copier.Copy(convertRest, req)
	return convertRest
}

func ConvertUPDtoToPo(req *user_dto.UserProfile) *user_po.UserProfile {
	return &user_po.UserProfile{
		AvatarUrl: req.AvatarUrl,
		Name:      req.Name,
		Email:     req.Email,
		Phone:     req.Phone,
		IsAdmin:   req.IsAdmin,
	}
}
