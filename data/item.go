package data

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
        Id          primitive.ObjectID `bson:"_id" json:"id"`
        Name        string             `bson:"name" json:"name"`
        Description string             `bson: "description" json:"description"`
        Price       float32            `bson:"price" json:"price"`
}
