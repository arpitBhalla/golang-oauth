package db

import (
	"context"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var clientInstance *mongo.Client

var clientInstanceError error

var mongoOnce sync.Once

const (
	DB    = "gawds"
	USERS = "users"
)

func GetMongoClient() (*mongo.Client, error) {
	connString := os.Getenv("MONGO_URI")

	if connString == "" {
		connString = "mongodb://localhost:27017"
	}
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(connString)
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			clientInstanceError = err
		}
		err = client.Ping(context.TODO(), nil)
		if err != nil {
			clientInstanceError = err
		}
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
