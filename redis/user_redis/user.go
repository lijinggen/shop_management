package user_redis

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/shop_management/sm_error"
	"github.com/shop_management/sm_error/error_code"
	"github.com/shop_management/vars"
	"time"
)

func getUserIdKey(userId string) string {
	return "session:user_id:" + userId
}

func getTokenKey(token string) string {
	return "session:token:" + token
}

func ClearSession(ctx *gin.Context, userId string) error {
	redisClient := vars.RedisClient
	tokenResp := redisClient.Get(ctx, getUserIdKey(userId))
	if tokenResp.Err() != nil {
		return nil
	}
	token := tokenResp.Val()
	err := redisClient.Del(ctx, getTokenKey(token)).Err()
	if err != nil {
		vars.Log.Errorf("ClearSession del token err:%v", err)
	}
	err = redisClient.Del(ctx, getUserIdKey(userId)).Err()
	if err != nil {
		vars.Log.Errorf("ClearSession del userId err:%v", err)
	}
	return nil
}

func SetSession(ctx *gin.Context, userId string, token string, expireTime int64) error {
	redisClient := vars.RedisClient
	err := redisClient.Set(ctx, getUserIdKey(userId), token, time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		vars.Log.Errorf("SetSession user_session_map err:%v, data:%s,%s", err, userId, token)
		return sm_error.NewHttpError(error_code.RedisErr)
	}
	err = redisClient.Set(ctx, getTokenKey(token), userId, time.Duration(expireTime)*time.Second).Err()
	if err != nil {
		vars.Log.Errorf("SetSession session_user_map err:%v, data:%s,%s", err, userId, token)
		return sm_error.NewHttpError(error_code.RedisErr)
	}
	return nil
}

func CheckToken(ctx *gin.Context, token string, userId string) error {
	get := vars.RedisClient.Get(ctx, getTokenKey(token))
	if get.Err() == redis.Nil || get.Val() == "" || get.Val() != userId {
		return sm_error.NewHttpError(error_code.UserTokenError)
	}
	return nil
}
