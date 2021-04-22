package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	uid, _ := r.Context().Value("id").(string)

	if uid == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Code:    401,
			Message: "Unauthorized",
		})
		return
	}

	client, err := db.GetMongoClient()

	var user UserData

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	docID, _ := primitive.ObjectIDFromHex(uid)

	res := collection.FindOne(context.TODO(), bson.D{bson.E{Key: "_id", Value: docID}})

	err = res.Decode(&user)

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "User Not Logged in",
		})
	} else {
		json.NewEncoder(w).Encode(UserResponse{
			Code:    200,
			Message: "Success",
			Result:  user,
		})
	}
}
