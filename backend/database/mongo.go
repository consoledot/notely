package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func Connect() {
	//  Set up connection string
	mongoURI := "mongodb://localhost:27017"

	// Set up options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Context creation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Mongo
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	Client = client
	log.Println("Connected to MongoDb")

}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}
	log.Println("Disconnected from MongoDb")
}
