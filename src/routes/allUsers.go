package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"gawds/src/models"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type AllUserResponse struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Result  []models.User `json:"result"`
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	client, err := db.GetMongoClient()

	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	col := client.Database(db.DB).Collection(db.USERS)

	cur, err := col.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var results []models.User

	for cur.Next(context.TODO()) {

		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(AllUserResponse{
		Code:    400,
		Message: "Connection Not Established",
		Result:  results,
	})

}
