package mongodb

import (
	"testing"

	"github.com/kittichai/core-framework/shared/framework/infrastructure/persistence/adapter/mongodb"
)

func TestMongoDBConnection(t *testing.T) {
	// Replace with your MongoDB URI
	uri := "mongodb://localhost:27017"

	db, err := mongodb.NewMongoDB(uri)
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.Close()

}
