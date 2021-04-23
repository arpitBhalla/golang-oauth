package routes

import (
	"context"
	"encoding/json"
	"gawds/src/db"
	"gawds/src/models"
	"gawds/src/utils"
	"net/http"

	"github.com/softbrewery/gojoi/pkg/joi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user LoginBody
	var loggedUser models.User

	client, connErr := db.GetMongoClient()

	emailError := joi.Validate(user.Email, joi.String().Email(nil))
	passwordError := joi.Validate(user.Password, joi.String().Min(6))

	if emailError != nil || passwordError != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body with error(s):" + emailError.Error() + passwordError.Error(),
		})
		return
	}

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

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: user.Email}})

	if res.Err() == nil {
		res.Decode(&loggedUser)

		err := bcrypt.CompareHashAndPassword([]byte(loggedUser.Password), []byte(user.Password))

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Login Failed",
			})
			return
		}

		tokens, err := utils.CreateToken(loggedUser.Uid.Hex())
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unable to create token",
			})
			return
		}
		redisError := utils.CreateAuth(loggedUser.Uid.Hex(), tokens)

		if redisError != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unable to save token",
			})
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
