package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type newUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser newUser

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body",
		})
		return
	}
	client, err := db.GetMongoClient()

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: newUser.Email}})

	if res.Err() == mongo.ErrNoDocuments {
		json.NewEncoder(w).Encode(Response{
			Code:    200,
			Message: "Success",
		})

		hash, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
		newUser.Password = string(hash)
		_, err = collection.InsertOne(context.TODO(), newUser)
		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Code:    400,
				Message: err.Error(),
			})
			return
		}

	} else {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "User Already Exists",
		})
		return
	}
}
