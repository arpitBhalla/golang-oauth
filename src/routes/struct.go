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
