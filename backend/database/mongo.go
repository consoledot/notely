package db

import (
	"context"
	"fmt"

	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ = godotenv.Load()
var (
	Client *mongo.Client
	dbName = os.Getenv("MONGO_DATABASE_NAME")
)

func Connect() {

	//  Set up connection string
	mongoURI := os.Getenv("MONGO_URI")
	fmt.Println(mongoURI)

	// Set up options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Context creation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to Mongo
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		log.Fatal("Error connecting", err)
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

func NotesCollection() *mongo.Collection {
	return Client.Database(dbName).Collection("notes")
}

func UsersCollection() *mongo.Collection {
	return Client.Database(dbName).Collection("users")
}
