package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client = Connection()

const DatabaseName = "potato"

func Connection() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

func OpenCollection(collectionName string) *mongo.Collection {
	collection := MongoClient.Database(DatabaseName).Collection(collectionName)
	return collection
}
