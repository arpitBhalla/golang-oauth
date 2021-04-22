package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
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
	var results []UserData

	for cur.Next(context.TODO()) {

		var elem UserData
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
		Code:   400,
		Result: results,
	})

}
