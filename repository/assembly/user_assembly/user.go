package user_assembly

import (
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/model"
)

func ConvertUDtoToModel(u *user_dto.User) *model.User {
	return &model.User{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		CreateTime: u.CreateTime,
		ModifyTime: u.ModifyTime,
		AvatarUrl:  u.AvatarUrl,
		Phone:      u.Phone,
	}
}

func ConvertUModelToDto(u *model.User) *user_dto.User {
	return &user_dto.User{
		Id:         u.Id,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		CreateTime: u.CreateTime,
		ModifyTime: u.ModifyTime,
		Phone:      u.Phone,
		AvatarUrl:  u.AvatarUrl,
	}
}
