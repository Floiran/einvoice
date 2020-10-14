package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	. "github.com/slovak-egov/einvoice/authproxy/config"
	"github.com/slovak-egov/einvoice/authproxy/user"
)

var ctx = context.Background()

type AuthDB interface {
	SaveUser(user *user.User)
	GetUserByToken(token string) *user.User
	GetUser(id string) *user.User

	AddToken(id, token string) error
	RemoveToken(id, token string) error
}

type redisDB struct {
	client *redis.Client
}

func NewAuthDB() AuthDB {
	rdb := redis.NewClient(&redis.Options{
		Addr:     Config.RedisUrl,
		Password: "", // no password set
		DB:       0,  // use default db
	})

	fmt.Println("ping", rdb.Ping(ctx).Val())

	return redisDB{rdb}
}

func (redisDB redisDB) SaveUser(user *user.User) {
	redisDB.client.Del(ctx, "user:"+user.Id)
	redisDB.client.HSet(ctx, "user:"+user.Id, "token", user.Token)
	redisDB.client.HSet(ctx, "user:"+user.Id, "name", user.Name)
	redisDB.client.HSet(ctx, "user:"+user.Id, "email", user.Email)
	redisDB.client.HSet(ctx, "user:"+user.Id, "serviceAccountKey", user.ServiceAccountKey)
}

func (redisDB redisDB) GetUserByToken(token string) *user.User {
	id := redisDB.client.HGet(ctx, "tokens", token).Val()
	if id == "" {
		return nil
	}
	return redisDB.GetUser(id)
}

func (redisDB redisDB) GetUser(id string) *user.User {
	if redisDB.client.Exists(ctx, "user:"+id).Val() == 0 {
		return nil
	}
	name := redisDB.client.HGet(ctx, "user:"+id, "name").Val()
	token := redisDB.client.HGet(ctx, "user:"+id, "token").Val()

	return &user.User{
		Token:             token,
		Id:                id,
		Name:              name,
		Email:             redisDB.client.HGet(ctx, "user:"+id, "email").Val(),
		ServiceAccountKey: redisDB.client.HGet(ctx, "user:"+id, "serviceAccountKey").Val(),
	}
}

func (redisDB redisDB) AddToken(id, token string) error {
	if redisDB.client.Exists(ctx, "user:"+id, token).Val() == 0 {
		return errors.New("User doesn't exist.")
	}
	redisDB.client.HSet(ctx, "user:"+id, "token", token)
	redisDB.client.HSet(ctx, "tokens", token, id)
	return nil
}

func (redisDB redisDB) RemoveToken(id, token string) error {
	if redisDB.client.Exists(ctx, "user:"+id, token).Val() == 0 {
		return errors.New("User doesn't exist.")
	}
	redisDB.client.HDel(ctx, "user:"+id, "token")
	redisDB.client.HDel(ctx, "tokens", token, id)
	return nil
}
