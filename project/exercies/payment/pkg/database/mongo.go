package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
Create a context
Connect to mongo
return the client instance
*/
type Database struct {
	Client *mongo.Client
	DB     *mongo.Database
}

func Connect() *Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		panic(err)
	}
	return &Database{
		Client: client,
		DB:     client.Database("wallet"),
	}
}
