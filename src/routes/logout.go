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

func Logout(w http.ResponseWriter, r *http.Request) {

	var (
		userData models.User
	)
	client, err := db.GetMongoClient()

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: userData.Email}})

	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.TODO(), userData)
			if err != nil {
				json.NewEncoder(w).Encode(Response{
					Code:    400,
					Message: "err",
				})
				return
			}
		} else {
			json.NewEncoder(w).Encode(Response{
				Code:    400,
				Message: "res.Err()",
			})
			return
		}
	}
}
