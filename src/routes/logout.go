package routes

import (
	"encoding/json"
	"gawds/src/utils"
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	metadata, err := utils.ExtractTokenMetadata(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Code:    401,
			Message: "Unauthorized",
		})
		return
	}
	delErr := utils.DeleteTokens(metadata)
	if delErr != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    500,
			Message: "Unable to log you out",
		})
		return
	}
	json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: "Done",
	})
}
