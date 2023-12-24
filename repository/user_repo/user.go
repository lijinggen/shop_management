package user_repo

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/model"
	"github.com/shop_management/repository"
	"github.com/shop_management/repository/assembly/user_assembly"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
	"github.com/shop_management/vars"
	"gorm.io/gorm"
	"time"
)

type userRepoImpl struct {
}

func NewUserRepoImpl() repository.UserRepo {
	return &userRepoImpl{}
}

func (u *userRepoImpl) Add(ctx *gin.Context, db *gorm.DB, dto *user_dto.User) error {
	err := db.Create(user_assembly.ConvertUDtoToModel(dto)).Error
	if err != nil {
		vars.Log.Errorf("userRepoImpl.Add error:%v,data: %v", err, util.MarshalToStringNoErr(dto))
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) GetById(ctx *gin.Context, db *gorm.DB, id string) (*user_dto.User, error) {
	user := &model.User{}
	err := db.Where("id = ?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		vars.Log.Errorf("userRepoImpl.GetById error:%v", err)
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	return user_assembly.ConvertUModelToDto(user), nil
}

func (u *userRepoImpl) GetBySecret(ctx *gin.Context, db *gorm.DB, secret string) (*user_dto.User, error) {
	user := &model.User{}
	err := db.Where("api_secret = ?", secret).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		vars.Log.Errorf("userRepoImpl.GetBySecret error:%v", err)
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	return user_assembly.ConvertUModelToDto(user), nil
}

func (u *userRepoImpl) GetByEmail(ctx *gin.Context, db *gorm.DB, email string) (*user_dto.User, error) {
	user := &model.User{}
	err := db.Where("email = ? ", email).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		vars.Log.Errorf("userRepoImpl.GetByEmail error:%v", err)
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	return user_assembly.ConvertUModelToDto(user), nil
}

func (u *userRepoImpl) GetByPhone(ctx *gin.Context, db *gorm.DB, phone string) (*user_dto.User, error) {
	user := &model.User{}
	err := db.Where("phone = ? ", phone).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		vars.Log.Errorf("userRepoImpl.GetByPhone error:%v", err)
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	return user_assembly.ConvertUModelToDto(user), nil
}

func (u *userRepoImpl) ModifyPassword(ctx *gin.Context, db *gorm.DB, id string, password string) error {
	err := db.Model(&model.User{}).Where("id=?", id).Update("password", password).Error
	if err != nil {
		vars.Log.Errorf("userRepoImpl.ModifyPassword error:%v", err)
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) SaveProfile(ctx *gin.Context, db *gorm.DB, req *user_dto.UserProfile) error {
	err := db.Model(&model.User{}).Where("id=?", req.UserId).Updates(map[string]interface {
	}{
		"phone":      req.Phone,
		"avatar_url": req.AvatarUrl,
		"name":       req.Name,
		"email":      req.Email,
	}).Error
	if err != nil {
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) Save(ctx *gin.Context, db *gorm.DB, req *user_dto.User) error {
	err := db.Save(user_assembly.ConvertUDtoToModel(req)).Error
	if err != nil {
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) Delete(ctx *gin.Context, db *gorm.DB, userId string) error {
	now := time.Now()
	err := db.Model(&model.User{}).Where("id = ?", userId).Update("deleted_time", &now).Error
	if err != nil {
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) SaveSecret(ctx *gin.Context, db *gorm.DB, userId, key, secret string) error {
	err := db.Model(&model.User{}).Where("id=?", userId).Updates(map[string]interface{}{
		"api_key":    key,
		"api_secret": secret,
	})
	if err != nil {
		vars.Log.Errorf("userRepoImpl.SaveSecret error:%v", err)
		return sm_error.NewHttpError(error_code.DBError)
	}
	return nil
}

func (u *userRepoImpl) List(ctx *gin.Context, db *gorm.DB, req *user_dto.UserListReq) ([]*user_dto.User, error) {
	if req.Name != "" {
		db = db.Where("name like ?", "%"+req.Name+"%")
	}
	if err := db.Model(&model.User{}).Count(&req.Pager.TotalRows).Error; err != nil {
		vars.Log.Errorf("userRepoImpl.List count error:%v,data: %v", err, util.MarshalToStringNoErr(req))
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	offset := (req.Pager.Page - 1) * req.Pager.PageSize

	mList := make([]*model.User, 0)
	err := db.Offset(int(offset)).Limit(int(req.Pager.PageSize)).Order("create_time desc").Find(&mList).Error
	if err != nil {
		vars.Log.Errorf("userRepoImpl.List Find error:%v,data: %v", err, util.MarshalToStringNoErr(req))
		return nil, sm_error.NewHttpError(error_code.DBError)
	}
	list := make([]*user_dto.User, 0)
	for _, user := range mList {
		list = append(list, user_assembly.ConvertUModelToDto(user))
	}
	return list, nil
}

func (u *userRepoImpl) GetAll(ctx *gin.Context, db *gorm.DB) ([]*user_dto.User, error) {
	mList := make([]*model.User, 0)
	err := db.Where("deleted_time is null").Order("create_time desc").Find(&mList).Error
	if err != nil {
		vars.Log.Errorf("userRepoImpl.GetAll Find error:%v", err)
		return nil, sm_error.NewHttpError(error_code.DBError)
	}

	list := make([]*user_dto.User, 0)
	for _, user := range mList {
		list = append(list, user_assembly.ConvertUModelToDto(user))
	}
	return list, nil
}
