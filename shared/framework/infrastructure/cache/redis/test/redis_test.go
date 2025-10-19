package test

import (
	"context"
	"testing"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/cache/redis"
)

func TestRedisConnection(t *testing.T) {

	db, err := redis.NewRedisDB("localhost:6379", "", 0)
	if err != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
	defer db.Close()
	if err := db.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("Failed to ping Redis: %v", err)
	}
}
