package routes

import (
	"encoding/json"
	"fmt"
	"gawds/src/utils"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func RefreshToken(w http.ResponseWriter, r *http.Request) {

	mapToken := map[string]string{}
	// if err := c.ShouldBindJSON(&mapToken); err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, err.Error())
	// 	return
	// }
	refreshToken := mapToken["refresh_token"]

	//verify the token
	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
	token, _ := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	//if there is an error, the token must have expired
	// if err != nil {
	// 	fmt.Println("the error: ", err)
	// 	c.JSON(http.StatusUnauthorized, "Refresh token expired")
	// 	return
	// }
	// //is token valid?
	// if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
	// 	c.JSON(http.StatusUnauthorized, err)
	// 	return
	// }
	//Since token is valid, get the uuid:
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
