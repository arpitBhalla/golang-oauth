package routes

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserData struct {
	Uid   primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
}
type UserResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Result  UserData `json:"result"`
}

type AllUserResponse struct {
	Code   int        `json:"code"`
	Result []UserData `json:"result"`
}

type LoginBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
}
