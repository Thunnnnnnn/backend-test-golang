package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection

func ConnectMongo() error {
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is empty")
	}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return err
	}

	UserCollection = client.Database("mongodb").Collection("users")

	return nil
}
