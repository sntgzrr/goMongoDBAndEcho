package repositories

import (
	"context"
	"echoApiRest/database"
	"echoApiRest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

// Creating collection users
var collection = database.SetCollection("users")

// Getting MongoDB context
var ctx = context.Background()

func Create(user models.User) error {
	var err error
	// Using InsertOne from MongoBD Driver to Create
	_, err = collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func Read() (models.Users, error) {
	var users models.Users
	// Using bson.D to decode the user
	filter := bson.D{}
	//	Using Find from MongoDB Driver to Read
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// Looping through cur to decode user and add it to users
	for cur.Next(ctx) {
		var user models.User
		// Decoding user
		err = cur.Decode(&user)
		if err != nil {
			return nil, err
		}
		// Adding user to users
		users = append(users, &user)
	}
	return users, nil
}

func Update(user models.User, userId string) error {
	var err error
	// Converting ID into PrimitiveID
	uid, _ := primitive.ObjectIDFromHex(userId)
	// Using bson.M for Specify field to update (ID)
	filter := bson.M{"_id": uid}
	// Using bson.M to Specify Fields to update in Mongo
	update := bson.M{
		"$set": bson.M{
			"name":       user.Name,
			"email":      user.Email,
			"updated_at": time.Now(),
		},
	}
	// Using UpdateOne from MongoDB Driver to update
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func Delete(userId string) error {
	var err error
	var uid primitive.ObjectID
	// Converting ID into PrimitiveID
	uid, err = primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}
	// Using bson.M for Specify document to delete by a field (ID)
	filter := bson.M{"_id": uid}
	// Using DeleteOne from MongoDB Driver to delete
	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
