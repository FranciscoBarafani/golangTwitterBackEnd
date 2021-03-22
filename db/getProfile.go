package db

import (
	"context"
	"fmt"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Gets user profile*/
func GetProfile(id string) (models.User, error) {
	/*Creates TimeOut Context to stop query if it lasts more than 15 seconds*/
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	/*Database Reference Creation*/
	db := MongoCN.Database("Twitter")
	col := db.Collection("Users")

	var profile models.User

	objectID, _ := primitive.ObjectIDFromHex(id)

	/*Query Condition*/
	condition := bson.M{
		"_id": objectID,
	}

	/*Query*/
	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		fmt.Println("Profile not found" + err.Error())
		return profile, err
	}
	return profile, nil
}
