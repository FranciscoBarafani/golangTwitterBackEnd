package db

import (
	"context"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Edit User Data */
func EditProfile(user models.User, id string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("Users")

	/* Makes an object with the data to update */
	data := make(map[string]interface{})
	if len(user.Name) > 0 {
		data["name"] = user.Name
	}
	if len(user.Lastname) > 0 {
		data["lastName"] = user.Lastname
	}
	if len(user.Avatar) > 0 {
		data["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		data["banner"] = user.Banner
	}
	if len(user.Bio) > 0 {
		data["bio"] = user.Bio
	}
	if len(user.WebSite) > 0 {
		data["webSite"] = user.WebSite
	}

	/* Update Query */
	updateString := bson.M{
		"$set": data,
	}

	/* Converts ID to ObjectID */
	objID, _ := primitive.ObjectIDFromHex(id)
	/* Creates ID filter */
	filter := bson.M{"_id": bson.M{"$eq": objID}}
	/* Update Query */
	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}
	return true, nil
}
