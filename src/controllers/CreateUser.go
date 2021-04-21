package controllers

import (
	"context"
	"gawds/src/db"
	"gawds/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(userData models.User) error {
	client, err := db.GetMongoClient()

	if err != nil {
		return err
	}

	collection := client.Database(db.DB).Collection(db.USERS)

	res := collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "uid", Value: userData.Uid}})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			_, err = collection.InsertOne(context.TODO(), userData)
			if err != nil {
				return err
			}
		} else {
			return res.Err()
		}
	} else {
		var user models.User
		err = res.Decode(&user)
		if err != nil {
			return err
		}
		filter := bson.D{primitive.E{Key: "uid", Value: userData.Uid}}
		updater := bson.D{
			primitive.E{Key: "$set", Value: userData},
		}
		_, err = collection.UpdateOne(context.TODO(), filter, updater)
		if err != nil {
			return err
		}
	}
	return nil
}
