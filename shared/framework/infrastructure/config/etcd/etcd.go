package etcd

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	logger "github.com/kittichai/core-framework/shared/framework/pkg/logger/applog"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var ErrMissingEndpoints = errors.New("missing ETCD_ENDPOINTS environment variable")

type EtcdClient struct {
	*clientv3.Client
	logger *logger.ECSLogger
}

type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
}

func NewEtcdClient(cfg Config, logger *logger.ECSLogger) (*EtcdClient, error) {
	ctx := logger.WithContext(context.Background())
	log := logger.FromContext(ctx)

	clientCfg := clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: cfg.DialTimeout,
	}
	client, err := clientv3.New(clientCfg)
	if err != nil {
		log.Error().
			Err(err).
			Strs("endpoints", cfg.Endpoints).
			Msg("Failed to create etcd client")
		return nil, err
	}

	// Verify connection
	ctx, cancel := context.WithTimeout(context.Background(), cfg.DialTimeout)
	defer cancel()
	_, err = client.Status(ctx, cfg.Endpoints[0])
	if err != nil {
		log.Error().
			Err(err).
			Str("endpoint", cfg.Endpoints[0]).
			Msg("Failed to ping etcd")
		client.Close()
		return nil, err
	}

	log.Info().
		Strs("endpoints", cfg.Endpoints).
		Msg("Successfully connected to etcd")
	return &EtcdClient{Client: client, logger: logger}, nil
}

func (c *EtcdClient) Close() error {
	ctx := c.logger.WithContext(context.Background())
	log := c.logger.FromContext(ctx)

	err := c.Client.Close()
	if err != nil {
		log.Error().
			Err(err).
			Msg("Failed to close etcd client")
		return err
	}
	log.Info().Msg("Etcd client closed")
	return nil
}

// LoadConfigFromEnv loads etcd configuration from environment variables
func LoadConfigFromEnv() (Config, error) {
	endpointsStr := os.Getenv("ETCD_ENDPOINTS")
	if endpointsStr == "" {
		return Config{}, ErrMissingEndpoints
	}
	endpoints := strings.Split(endpointsStr, ",")

	dialTimeoutStr := os.Getenv("ETCD_DIAL_TIMEOUT")
	if dialTimeoutStr == "" {
		dialTimeoutStr = "5s" // Default value
	}
	dialTimeout, err := time.ParseDuration(dialTimeoutStr)
	if err != nil {
		return Config{}, err
	}

	return Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	}, nil
}
