package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"net/http"

	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type loginBody struct {
	Email    int    `json:"email"`
	Password string `json:"password"`
}

var store = sessions.NewCookieStore([]byte("secretKey"))

func Login(w http.ResponseWriter, r *http.Request) {
	var user loginBody

	client, connErr := db.GetMongoClient()

	if connErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Connection Not Established",
		})
		return
	}

	jsonErr := json.NewDecoder(r.Body).Decode(&user)
	if jsonErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body",
		})
		return
	}

	session, _ := store.Get(r, "session-name")
	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}})

	if res.Err() != nil {
		session.Values["auth-token"] = ""
		err := session.Save(r, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
