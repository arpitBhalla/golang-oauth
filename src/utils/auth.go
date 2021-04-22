package utils

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

func ExtractToken(r *http.Request) (*jwt.Token, error) {

	bearToken := r.Header.Get("Authorization")

	tokenString := strings.Split(bearToken, " ")

	tokenValue := ""
	if len(tokenString) == 2 {
		tokenValue = tokenString[1]

		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return "ACCESS_SECRET", nil
		})
		if err != nil {
			return nil, err
		}
		return token, nil
	} else {
		return nil, errors.New("Unauthenticated")
	}
}

func CreateToken(userid string) (*Tokens, error) {
	tokenData := &Tokens{}
	tokenData.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	tokenData.AccessUuid = uuid.NewV4().String()

	tokenData.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	tokenData.RefreshUuid = uuid.NewV4().String()
	print(80)

	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = tokenData.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = tokenData.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenData.AccessToken, err = at.SignedString([]byte("accessSecret"))
	if err != nil {
		print(err.Error())
		return nil, err
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = tokenData.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = tokenData.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	tokenData.RefreshToken, err = rt.SignedString([]byte("tokenSecret"))
	if err != nil {
		return nil, err
	}
	return tokenData, nil
}

func CreateAuth(userid string, td *Tokens) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(td.AccessUuid, (userid), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := client.Set(td.RefreshUuid, (userid), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}
func DeleteAuth(givenUuid string) (int64, error) {
	deleted, err := client.Del(givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}
