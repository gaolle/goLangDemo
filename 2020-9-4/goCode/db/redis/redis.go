package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// redis
var (
	redisClient *redis.Client
	redisOptions redis.Options = redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	}
)

// 初始化客户端
func init() {
	redisClient = redis.NewClient(&redisOptions)
	pong, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
	if pong != "PONG" {
		fmt.Println("pong failed")
	} else {
		fmt.Println("redis pong")
	}
}

// 设置键值
func Set(key string, value string, expiration time.Duration) error {
	return redisClient.Set(key, value, expiration).Err()
}

// 删除操作 删除成功返回值为nil,删除不存在的返回值也为nil
func Del(keys ...string) error {
	if len(keys) == 0 {
		return errors.New("no value")
	}
	if len(keys) == 1 {
		return redisClient.Del(keys...).Err()
	} else {
		return redisClient.Del(keys...).Err()
	}
}

// 查询键值
func Get(key string) (string, error) {
	val, err := redisClient.Get(key).Result()
	return val, err
}

// 修改旧值并返回新值
func GetSet(key string, value string) (string, error) {
	val, err := redisClient.GetSet(key, value).Result()
	return val, err
}