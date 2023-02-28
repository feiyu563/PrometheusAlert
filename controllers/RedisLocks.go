package controllers

import (
	"context"
	"github.com/astaxie/beego"
	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
	"time"
)

var redisClient *redis.Client

func InitRedisLockClient() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     beego.AppConfig.String("redis_addr"),
		Username: beego.AppConfig.String("redis_user"),
		Password: beego.AppConfig.String("redis_password"),
	})
}

func NewRedisLock(ctx context.Context, key string, timeOut time.Duration, retryInterval time.Duration, tryMax int) (*redislock.Lock, error) {
	locker := redislock.New(redisClient)

	// 在 TryMax 重试次数中，周期尝试获取锁
	backoff := redislock.LimitRetry(redislock.LinearBackoff(retryInterval), tryMax)
	lock, err := locker.Obtain(ctx, key, timeOut, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err == redislock.ErrNotObtained {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return lock, nil
}
