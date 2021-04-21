package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"gawds/src/db"
	"gawds/src/models"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	uid, _ := r.Context().Value("id").(string)

	// w.WriteHeader(http.StatusOK)

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

	docID, _ := primitive.ObjectIDFromHex(uid)

	res := collection.FindOne(context.TODO(), bson.M{"_id": docID})

	err = res.Decode(&user)

	fmt.Println(user, err, (uid))

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
