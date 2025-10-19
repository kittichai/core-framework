// shared/pkg/config/config.go
package config

import (
	"github.com/kittichai/core-framework/shared/framework/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Postgres struct{ DSN string }
		Mongo    struct{ URI string }
	}
	Redis struct{ Addr string }
	Kafka struct{ Brokers []string }
	Etcd  struct{ Endpoints []string }
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.NewAppError(500, "Failed to load config", err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.NewAppError(500, "Failed to parse config", err)
	}
	return &cfg, nil
}
