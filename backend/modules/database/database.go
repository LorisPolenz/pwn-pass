package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetMongodbContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	return ctx
}

func GetMongodbClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	client.Connect(GetMongodbContext())

	if err != nil {
		log.Fatal(err)
	}

	return client
}
