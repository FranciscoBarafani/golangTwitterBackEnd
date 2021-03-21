package db

import (
	"context"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

/* Check if user already exists */
func CheckUserExists(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter")
	collection := db.Collection("Users")

	condition := bson.M{"email": email}

	var result models.User

	err := collection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
