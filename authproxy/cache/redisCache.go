package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
	"github.com/slovak-egov/einvoice/authproxy/config"
)

var ctx = context.Background()

type redisCache struct {
	client *redis.Client
}

func New() Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Config.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default db
	})

	ping := rdb.Ping(ctx).Val()
	if ping == "" {
		log.WithField("redis_url", config.Config.RedisUrl).Error("app.redis.not_connected")
	} else {
		log.Info("app.redis.connected")
	}

	return &redisCache{rdb}
}

func userIdKey(token string) string {
	return "token-" + token
}

func (redis *redisCache) SaveToken(token, userId string) {
	redis.client.Set(ctx, userIdKey(token), userId, config.Config.TokenExpiration).Val()
}

func (redis *redisCache) GetUserId(token string) (string, error) {
	id := redis.client.Get(ctx, userIdKey(token)).Val()
	if id == "" {
		log.WithField("token", token).Debug("redis.token.not_found")
		return "", errors.New("Token not found")
	}
	redis.client.Expire(ctx, userIdKey(token), config.Config.TokenExpiration)
	return id, nil
}

func (redis *redisCache) RemoveToken(token string) bool {
	deletedRecordsNumber := redis.client.Del(ctx, userIdKey(token)).Val()
	return deletedRecordsNumber == 1
}
