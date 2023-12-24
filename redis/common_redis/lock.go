package common_redis

import _ "github.com/redis/go-redis/v9"

func Lock() (bool, error) {
	return false, nil
}

func UnLock() {

}
