package routes

import (
	"encoding/json"
	"gawds/src/controllers"
	"gawds/src/models"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type userResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  models.User `json:"result"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	err = controllers.CreateUser(newUser)
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(response{
			Code:    200,
			Message: "Success",
		})
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	user, err := controllers.GetUser(vars["uid"])

	if err != nil {
		json.NewEncoder(w).Encode(response{
			Code:    400,
			Message: err.Error(),
		})
	} else {
		json.NewEncoder(w).Encode(userResponse{
			Code:    200,
			Message: "Success",
			Result:  user,
		})
	}
}
