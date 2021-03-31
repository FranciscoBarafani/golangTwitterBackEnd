package db

import (
	"context"
	"log"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTweet(ID string, page int64) ([]*models.GetTweet, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("Twitter")
	col := db.Collection("Tweets")

	var results []*models.GetTweet

	/* Query Condition */
	condition := bson.M{
		"userId": ID,
	}

	/* Find Options */
	options := options.Find()
	/* Sets result limit to 20 */
	options.SetLimit(20)
	/* Order by date, Value -1 sets descendent order */
	options.SetSort(bson.D{{Key: "date", Value: -1}})
	/* Set how many records to skip for page changes */
	options.SetSkip((page - 1) * 20)

	result, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for result.Next(context.TODO()) {
		/* For each iteration it creates a new record, decodes it and adds it
		to the final result */
		var record models.GetTweet
		err := result.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
