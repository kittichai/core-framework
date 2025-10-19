package infrastructure

import (
	"context"
	"log"

	etcd "github.com/kittichai/core-framework/shared/framework/infrastructure/config/etcd"
	logger "github.com/kittichai/core-framework/shared/framework/pkg/logger/applog"
)

func InitConnection() (ctx Context, closeConnection func(), err error) {

	ctx = Context{}

	ctx.logger = logger.NewECSLogger()

	// Initialize all database connections here

	//initialize ETCD
	cfg, err := etcd.LoadConfigFromEnv()
	if err != nil {
		ctx.logger.Error(context.Background()).
			Err(err).
			Msg("Failed to load etcd config from env")
		log.Fatalf("failed to load etcd config: %v", err)
	}

	ctx.etcd, err = etcd.NewEtcdClient(cfg, ctx.logger)
	if err != nil {
		return ctx, closeConnection, err
	}

	closeConnection = func() {
		//close all connections here
		if ctx.redis != nil {
			ctx.redis.Close()
		}
		if ctx.postgres != nil {
			ctx.postgres.Close()
		}
		if ctx.mysql != nil {
			ctx.mysql.Close()
		}
		if ctx.mongodb != nil {
			ctx.mongodb.Close()
		}
		if ctx.rabbitmq != nil {
			ctx.rabbitmq.Close()
		}
		if ctx.kafka != nil {
			ctx.kafka.Close()
		}
		if ctx.cassandra != nil {
			ctx.cassandra.Close()
		}
		if ctx.etcd != nil {
			ctx.etcd.Close()
		}
	}
	return ctx, closeConnection, nil
}
