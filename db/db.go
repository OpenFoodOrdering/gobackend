package db

import (
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb mongo.Client

func Init(url string) {

	// Setup MongoDb Options
	ServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	ClientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(ServerAPI)

	// Initialize MongoDb Client Connection Pool Using ClientOptions
	MongoDb, err := mongo.NewClient(ClientOptions)
	if err != nil {
		log.Fatal(err)
	}
}
