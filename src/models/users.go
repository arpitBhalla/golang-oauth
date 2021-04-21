package models

type User struct {
	Uid  string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
