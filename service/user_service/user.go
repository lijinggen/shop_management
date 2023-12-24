package user_service

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/shop_management/dto/user_dto"
	"github.com/shop_management/redis/user_redis"
	"github.com/shop_management/repository"
	"github.com/shop_management/repository/user_repo"
	"github.com/shop_management/service"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/util"
	"math/rand"
)

type userServiceImpl struct {
	userRepo repository.UserRepo
}

func NewUserServiceImpl() service.UserService {
	return &userServiceImpl{
		userRepo: user_repo.NewUserRepoImpl(),
	}
}

const tokenExpireTime = 3600 * 24 * 3

func (u *userServiceImpl) Register(ctx *gin.Context, req *user_dto.RegisterUserReq) error {
	var err error
	tx := util.GetDBFromContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	phone, err := u.userRepo.GetByPhone(ctx, tx, req.Phone)
	if err != nil {
		return err
	}
	if phone != nil {
		return sm_error.NewHttpError(error_code.UserPhoneExists)
	}
	if req.Password != req.ConfirmPassword {
		return sm_error.NewHttpError(error_code.UserConfirmPasswordIncorrect)
	}
	err = u.userRepo.Add(ctx, tx, &user_dto.User{
		Id:       uuid.New().String(),
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return err
	}
	return nil
}

func (u *userServiceImpl) Login(ctx *gin.Context, req *user_dto.UserLogin) (string, error) {
	tx := util.GetDBFromContext(ctx).Begin()
	accountDetail, err := u.userRepo.GetByPhone(ctx, tx, req.Phone)
	if err != nil {
		return "", err
	}
	if accountDetail == nil {
		return "", sm_error.NewHttpError(error_code.UserLoginFailed, "Email not exists!")
	}
	if accountDetail.Password == req.Password {
		token := generateRandomToken()
		// 设置cookie
		ctx.SetCookie("token", token, tokenExpireTime, "/", "", false, false)
		ctx.SetCookie("user_id", accountDetail.Id, tokenExpireTime, "/", "", false, false)
		// 设置session到redis
		user_redis.ClearSession(ctx, accountDetail.Id)
		user_redis.SetSession(ctx, accountDetail.Id, token, tokenExpireTime)
		tx.Commit()
		return accountDetail.Id, nil
	} else {
		return "", sm_error.NewHttpError(error_code.UserLoginFailed)
	}
}

//	func (u *userServiceImpl) ModifyPwd(ctx *gin.Context, req *user_dto.ModifyUserPasswordReq) error {
//		tx := util.GetDBFromContext(ctx).Begin()
//		var err error
//		defer func() {
//			if err != nil {
//				tx.Rollback()
//			} else {
//				tx.Commit()
//			}
//		}()
//		accountDetail, err := u.userRepo.GetById(ctx, tx, req.UserId)
//		if err != nil {
//			return err
//		}
//		if accountDetail.Password != req.OldPassword {
//			return js_error.NewHttpError(error_code.UserModifyPasswordFailed)
//		}
//		if req.NewPassword != req.ConfirmPassword {
//			return js_error.NewHttpError(error_code.UserModifyPasswordFailed)
//		}
//		err = u.userRepo.ModifyPassword(ctx, tx, accountDetail.Id, req.NewPassword)
//		if err != nil {
//			return err
//		}
//		_ = user_redis.ClearSession(ctx, accountDetail.Id)
//		return nil
//	}
func (u *userServiceImpl) GetUserProfile(ctx *gin.Context, userId string) (*user_dto.UserProfile, error) {
	user, err := u.userRepo.GetById(ctx, util.GetDBFromContext(ctx), userId)
	if err != nil {
		return nil, sm_error.NewHttpError(error_code.UserModifyPasswordFailed)
	}
	if user == nil {
		return nil, sm_error.NewHttpError(error_code.UserNoExists)
	}
	return &user_dto.UserProfile{
		AvatarUrl: user.AvatarUrl,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
	}, nil
}

//	func (u *userServiceImpl) SaveUserProfile(ctx *gin.Context, req *user_dto.UserProfile) error {
//		user, err := u.userRepo.GetById(ctx, util.GetDBFromContext(ctx), req.UserId)
//		if err != nil {
//			return err
//		}
//		if user == nil {
//			return js_error.NewHttpError(error_code.UserNoExists)
//		}
//		err = u.userRepo.SaveProfile(ctx, util.GetDBFromContext(ctx), req)
//		if err != nil {
//			return err
//		}
//		return nil
//	}
//
//	func (u *userServiceImpl) Deactivate(ctx *gin.Context, userId string) error {
//		err := u.userRepo.Delete(ctx, util.GetDBFromContext(ctx), userId)
//		if err != nil {
//			return err
//		}
//		_ = user_redis.ClearSession(ctx, userId)
//		return nil
//	}
func generateRandomToken() string {
	bytes := make([]byte, 24)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	sessionValue := hex.EncodeToString(bytes)
	return sessionValue
}

//
//func (u *userServiceImpl) GenerateApiSecret(ctx *gin.Context, req *user_dto.GenerateApiSecret) error {
//	user, err := u.userRepo.GetById(ctx, util.GetDBFromContext(ctx), req.UserId)
//	if err != nil {
//		return err
//	}
//	if user == nil {
//		return js_error.NewHttpError(error_code.UserNoExists)
//	}
//	apiSecret := generateRandomToken()
//	err = u.userRepo.SaveSecret(ctx, util.GetDBFromContext(ctx), req.UserId, req.ApiKey, apiSecret)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *userServiceImpl) GetApiSecret(ctx *gin.Context, userId string) (*user_dto.GetApiSecret, error) {
//	user, err := u.userRepo.GetById(ctx, util.GetDBFromContext(ctx), userId)
//	if err != nil {
//		return nil, err
//	}
//	if user == nil {
//		return nil, js_error.NewHttpError(error_code.UserNoExists)
//	}
//	return &user_dto.GetApiSecret{
//		ApiKey:    user.ApiKey,
//		ApiSecret: user.ApiSecret,
//	}, nil
//}
//
//func (u *userServiceImpl) UserList(ctx *gin.Context, req *user_dto.UserListReq) (*user_dto.UserListResp, error) {
//	list, err := u.userRepo.List(ctx, util.GetDBFromContext(ctx), req)
//	if err != nil {
//		return nil, err
//	}
//	ids := make([]string, len(list))
//	for i, user := range list {
//		ids[i] = user.Id
//	}
//	balanceMap, err := u.userBalanceService.BatchGetUserBalance(ctx, ids)
//	if err != nil {
//		return nil, err
//	}
//	for _, user := range list {
//		if v, ok := balanceMap[user.Id]; ok {
//			user.Balance = v.Balance
//			user.Currency = v.Currency
//		}
//	}
//	return &user_dto.UserListResp{
//		Pager: req.Pager,
//		List:  list,
//	}, nil
//}
//
//func (u *userServiceImpl) SaveUser(ctx *gin.Context, req *user_dto.User) error {
//	user, err := u.userRepo.GetById(ctx, util.GetDBFromContext(ctx), req.Id)
//	if err != nil {
//		return err
//	}
//	if user == nil {
//		return js_error.NewHttpError(error_code.UserNoExists)
//	}
//	ub, err := u.userBalanceService.GetByUserId(ctx, req.Id)
//	if err != nil {
//		return err
//	}
//	if ub != nil {
//		ub.Balance = req.Balance
//		ub.Currency = req.Currency
//	} else {
//		ub = &user_dto.UserBalance{
//			UserId:   req.Id,
//			Balance:  req.Balance,
//			Currency: req.Currency,
//		}
//	}
//
//	user.Name = req.Name
//	user.Email = req.Name
//	user.Password = req.Password
//	user.PhoneNumber = req.PhoneNumber
//	err = u.userRepo.Save(ctx, util.GetDBFromContext(ctx), user)
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *userServiceImpl) GetUserIdBySecret(ctx *gin.Context, key string, secret string) (string, error) {
//	user, err := u.userRepo.GetBySecret(ctx, util.GetDBFromContext(ctx), secret)
//	if err != nil {
//		return "", err
//	}
//	if user == nil {
//		return "", js_error.NewHttpError(error_code.UserSecretIncorrect)
//	}
//	if user.ApiKey != key {
//		return "", js_error.NewHttpError(error_code.UserSecretIncorrect)
//	}
//	return user.Id, nil
//}
