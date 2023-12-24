package vars

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var DbMetadataName = "joinner_db"
var Log *zap.SugaredLogger
var RedisClient *redis.Client

//var ActionAuthService *action_auth.ActionAuthService
