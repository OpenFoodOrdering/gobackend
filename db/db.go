package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

	ctx, cancelfunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelfunc()
	// Ping MongoDb
	err = MongoDb.Ping(ctx, &readpref.ReadPref{})
	if err != nil {
		log.Fatal("Unable to Connect to Mongo Db:", err)
	}
}

func GetClient() *mongo.Client {
	return &MongoDb
}
