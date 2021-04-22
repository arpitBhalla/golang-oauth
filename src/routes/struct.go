package routes

import (
	"gawds/src/models"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type UserResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  models.User `json:"result"`
}

type AllUserResponse struct {
	Code   int           `json:"code"`
	Result []models.User `json:"result"`
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
