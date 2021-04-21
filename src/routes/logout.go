package routes

import (
	"encoding/json"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session-name")

	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Unable to logout",
		})
	} else {
		json.NewEncoder(w).Encode(Response{
			Code:    200,
			Message: "Done",
		})
	}
}
