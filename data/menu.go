package data

import (
        "go.mongodb.org/mongo-driver/bson/primitive"
        "time"
)

type Menu struct {
        Id            primitive.ObjectID `bson:"_id" json:"_id"`
        Title         string             `bson:"title" json:"title"`
        Offered_From  time.Time          `bson:"start_time" json:"start_time"`
        Offered_Until time.Time          `bson:"end_time" json:"end_time"`
		Items 		  []Item 			 `bson:"items" json:"items"`
}
