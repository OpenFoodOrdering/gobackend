package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDb mongo.Client

func Init(url string) {

	// Setup MongoDb Options
	ServerAPI := options.ServerAPI(options.ServerAPIVersion1)
	ClientOptions := options.Client().ApplyURI(url).SetServerAPIOptions(ServerAPI)

	ctx := context.Background()
	// Initialize MongoDb Client Connection Pool Using ClientOptions
	MongoDb, err := mongo.Connect(ctx, ClientOptions)
	if err != nil {
		log.Fatal("Mongodb_Initial_Connection: ", err)
	}

	// Ping Mongo
	if err = MongoDb.Database("openfoodordering").RunCommand(ctx, bson.D{primitive.E{Key: "ping", Value: 1}}).Err(); err != nil {
		log.Fatal("Mongodb:", err)
	}

	log.Println("MongoDB Connection Successful")
}

func GetClient() *mongo.Client {
	return &MongoDb
}
