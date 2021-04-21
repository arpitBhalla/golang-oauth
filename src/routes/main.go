package routes

import (
	"encoding/json"
	"gawds/src/models"
	"net/http"
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

func newUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		json.NewEncoder(w).Encode(response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	err = CreateUser(newUser)
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

func getUser(w http.ResponseWriter, r *http.Request, uid string) {
	user, err := GetUser(uid)

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