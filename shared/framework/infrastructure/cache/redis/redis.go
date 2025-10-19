package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisDB struct {
	*redis.Client
}

func NewRedisDB(addr, password string, db int) (*RedisDB, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return &RedisDB{Client: rdb}, nil
}

func (db *RedisDB) Close() error {
	return db.Client.Close()
}
