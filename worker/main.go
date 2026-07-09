package main

import (
	"backend-test-golang/database"
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	if err := database.ConnectMongo(); err != nil {
		log.Fatal("Mongo connection failed:", err)
	}
	for {
		ctx := context.Background()

		if database.UserCollection == nil {
			log.Println("No user collection")

		} else {

			cursor, err := database.UserCollection.Find(ctx, bson.M{})

			if err != nil {
				log.Println("Error finding users:", err)
			} else {
				var users []interface{}
				if err := cursor.All(ctx, &users); err != nil {
					log.Println("Error decoding users:", err)
				} else {
					log.Printf("Current users %d\n", len(users))
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}
