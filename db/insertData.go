package db

import (
	"context"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Insert User Data */
func InsertData(user models.User) (string, bool, error) {
	/*Setting TimeOut Time */
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	/*Defer is called last in the function */
	defer cancel()

	db := MongoCN.Database("Twitter")
	collection := db.Collection("Users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
