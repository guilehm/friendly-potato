package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

func OpenCollection(collectionName, databaseName string) *mongo.Collection {
	if databaseName == "" {
		databaseName = DatabaseName
	}
	collection := MongoClient.Database(databaseName).Collection(collectionName)
	return collection
}

func createIndex(key string, unique bool, collection *mongo.Collection) {
	indexName, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.M{
				key: 1,
			},
			Options: options.Index().SetUnique(unique),
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"Successfully created index %s for collection %s\n", indexName, collection.Name(),
	)
}

func CreateIndexes() {
	gamesCollection := OpenCollection("games", "")
	usersCollection := OpenCollection("users", "")
	unSitemapsCollection := OpenCollection("sitemaps", "un")

	createIndex("id", true, gamesCollection)
	createIndex("slug", true, gamesCollection)

	createIndex("email", true, usersCollection)
	createIndex("token", false, usersCollection)
	createIndex("refresh_token", false, usersCollection)

	createIndex("location", true, unSitemapsCollection)
	createIndex("crawled", false, unSitemapsCollection)
}
