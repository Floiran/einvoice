package cache

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

type redisCache struct {
	userTokenExpiration time.Duration
	client *redis.Client
}

func NewRedis(redisUrl string, userTokenExpiration time.Duration) Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: "", // no password set
		DB:       0,  // use default db
	})

	pong := rdb.Ping(ctx).Val()
	if pong == "" {
		log.WithField("redis_url", redisUrl).Fatal("app.redis.not_connected")
	} else {
		log.Info("app.redis.connected")
	}

	return &redisCache{
		userTokenExpiration: userTokenExpiration,
		client: rdb,
	}
}

func userIdKey(token string) string {
	return "token-" + token
}

func (redis *redisCache) SaveUserToken(token, userId string) {
	redis.client.Set(ctx, userIdKey(token), userId, redis.userTokenExpiration).Val()
}

func (redis *redisCache) GetUserId(token string) (string, error) {
	id := redis.client.Get(ctx, userIdKey(token)).Val()
	if id == "" {
		log.WithField("token", token).Debug("redis.token.not_found")
		return "", errors.New("Token not found")
	}
	redis.client.Expire(ctx, userIdKey(token), redis.userTokenExpiration)
	return id, nil
}

func (redis *redisCache) RemoveUserToken(token string) bool {
	return redis.client.Del(ctx, userIdKey(token)).Val() == 1
}
