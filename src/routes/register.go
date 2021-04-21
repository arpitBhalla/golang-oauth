package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"gawds/src/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: newUser.Email}})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.TODO(), newUser)
			if err != nil {
				return err
			}
		} else {
			return res.Err()
		}
	} else {
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Code:    400,
				Message: err.Error(),
			})
		} else {
			json.NewEncoder(w).Encode(Response{
				Code:    200,
				Message: "Success",
			})
		}
	}
}
