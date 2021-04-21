package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"gawds/src/models"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	client, err := db.GetMongoClient()

	var user models.User

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	filter := bson.D{bson.E{Key: "_id", Value: vars["uid"]}}

	res := collection.FindOne(context.TODO(), filter)

	err = res.Decode(&user)

	// return user, err

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "User Not Found",
		})
	} else {
		json.NewEncoder(w).Encode(UserResponse{
			Code:    200,
			Message: "Success",
			Result:  user,
		})
	}
}
