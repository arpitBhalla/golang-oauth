package routes

import (
	"encoding/json"
	"fmt"
	"gawds/src/utils"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	RefreshToken string `json:"refreshToken"`
}

func RefreshToken(w http.ResponseWriter, r *http.Request) {
	var bodyToken Token

	err := json.NewDecoder(r.Body).Decode(&bodyToken)
	if err != nil {
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Invalid Body",
		})
		return
	}
	refreshToken := bodyToken.RefreshToken

	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf")
	token, _ := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		fmt.Println("the error: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: "Refresh token expired",
		})
		return
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {

			w.WriteHeader(http.StatusUnprocessableEntity)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unauthorized",
			})
			return
		}
		userId := claims["user_id"].(string)

		deleted, delErr := utils.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 { //if any goes wrong
			w.WriteHeader(http.StatusUnauthorized)

			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unauthorized",
			})
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := utils.CreateToken(userId)
		if createErr != nil {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unable to save token",
			})
			return
		}
		//save the tokens metadata to redis
		saveErr := utils.CreateAuth(userId, ts)
		if saveErr != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"code":    400,
				"Message": "Unable to save token",
			})
			return
		}
		tokens := map[string]interface{}{
			"code":         200,
			"accessToken":  ts.AccessToken,
			"refreshToken": ts.RefreshToken,
		}
		json.NewEncoder(w).Encode(tokens)
	} else {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"code":    400,
			"Message": "Refresh Expired",
		})
	}
}
