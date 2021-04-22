package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"gawds/src/models"
	"gawds/src/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user LoginBody
	var loggedUser models.User

	client, connErr := db.GetMongoClient()

	if connErr != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	jsonErr := json.NewDecoder(r.Body).Decode(&user)

	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body",
		})
		return
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}, primitive.E{Key: "password", Value: user.Password}})

	if res.Err() == nil {
		res.Decode(&loggedUser)
		tokens, err := utils.CreateToken(loggedUser.Uid.Hex())
		err2 := utils.CreateAuth(loggedUser.Uid.Hex(), tokens)

		if err != nil || err2 != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(LoginResponse{
			Code:         200,
			Message:      "Login Done",
			RefreshToken: tokens.RefreshToken,
			AccessToken:  tokens.AccessToken,
		})
	} else {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Login Failed",
		})
	}
}
