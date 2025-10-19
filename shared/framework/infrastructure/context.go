package infrastructure

import (
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/cache/redis"
	etcd "github.com/kittichai/core-framework/shared/framework/infrastructure/config/etcd"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/messaging/kafka"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/messaging/rabbitmq"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/cassandra"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/mongodb"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/mysql"
	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/postgres"
	logger "github.com/kittichai/core-framework/shared/framework/pkg/logger/applog"
)

type Context struct {
	etcd        *etcd.EtcdClient
	redis       *redis.RedisDB
	postgres    *postgres.PostgresDB
	mysql       *mysql.MySQLDB
	mongodb     *mongodb.MongoDB
	rabbitmq    *rabbitmq.RabbitMQConn
	kafka       *kafka.KafkaConnection
	cassandra   *cassandra.CassandraDB
	logger      *logger.ECSLogger
	healthCheck *healthcheck.HealthChecker
}
