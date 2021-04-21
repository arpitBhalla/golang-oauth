package routes

import (
	"context"
	"gawds/src/db"
	"gawds/src/models"

	"go.mongodb.org/mongo-driver/bson"
)

func GetUser(uid string) (models.User, error) {
	client, err := db.GetMongoClient()

	var user models.User

	if err != nil {
		return user, err
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	filter := bson.D{bson.E{Key: "uid", Value: uid}}

	res := collection.FindOne(context.TODO(), filter)

	err = res.Decode(&user)

	return user, err
}
