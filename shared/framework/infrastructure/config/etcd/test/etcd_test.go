package tests

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/config/etcd"
	logger "github.com/kittichai/core-framework/shared/framework/pkg/logger/applog"
)

func TestEtcdClient(t *testing.T) {
	// Use a real etcd instance for integration testing
	// For unit testing, consider using a mock or testcontainers

	logger := logger.NewECSLogger()
	//initialize ETCD
	cfg, err := etcd.LoadConfigFromEnv()
	if err != nil {
		logger.Error(context.Background()).
			Err(err).
			Msg("Failed to load etcd config from env")
		log.Fatalf("failed to load etcd config: %v", err)
	}
	client, err := etcd.NewEtcdClient(cfg, logger)
	if err != nil {
		t.Fatalf("failed to connect to etcd: %v", err)
	}
	defer client.Close()

	t.Run("Ping Etcd", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()
		_, err := client.Status(ctx, "localhost:2379")
		if err != nil {
			t.Fatalf("failed to ping etcd: %v", err)
		}
	})
}
