package tests

import (
	"testing"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/cassandra"
)

func TestCassandraConnection(t *testing.T) {
	db, err := cassandra.NewCassandraDB([]string{"localhost:9042"}, "test_keyspace")
	if err != nil {
		t.Fatalf("failed to connect to Cassandra: %v", err)
	}
	defer db.Close()
}
