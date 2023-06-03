package data

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/OpenFoodOrdering/gobackend/db"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Menu struct {
	Id            primitive.ObjectID `bson:"_id" json:"_id"`
	Title         string             `bson:"title" json:"title"`
	Offered_From  time.Time          `bson:"start_time" json:"start_time"`
	Offered_Until time.Time          `bson:"end_time" json:"end_time"`
	Items         []Item             `bson:"items" json:"items"`
}

// Handler that Gets a Single Menu
func GetOneMenu(w http.ResponseWriter, r *http.Request) {
	// Set Content Type To json
	w.Header().Set("Content-Type", "application/json")

	// Get a 10 Second Context
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	// Defer Context Handler
	defer cancel()
	// Database Collection
	collection := db.GetClient().Database("openfoodordering").Collection("menus")

	var result Menu
	id_val := chi.URLParamFromCtx(ctx, "id")
	id, err := primitive.ObjectIDFromHex(id_val)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get Filter to Parse Url Params
	filter := bson.D{primitive.E{Key: "_id", Value: id}}

	err = collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err != nil {
		// Todo Do Better Error Codes
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cancel()
	// Send The Result Away
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
