package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"net/http"

	"github.com/softbrewery/gojoi/pkg/joi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type newUser struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
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

	nameError := joi.Validate(newUser.Email, joi.String().Email(nil))
	emailError := joi.Validate(newUser.Email, joi.String().Email(nil))
	passwordError := joi.Validate(newUser.Password, joi.String().Min(6))

	if emailError != nil || passwordError != nil || nameError != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body with error(s):" + emailError.Error() + passwordError.Error() + nameError.Error(),
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
		json.NewEncoder(w).Encode(Response{
			Code:    200,
			Message: "Successfully registered",
		})

	} else {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "User Already Exists",
		})
		return
	}
}
