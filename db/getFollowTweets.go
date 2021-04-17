package db

import (
	"context"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"go.mongodb.org/mongo-driver/bson"
)

func GetFollowTweets(ID string, page int) ([]models.ReturnFollowTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("Relation")

	skip := ((page - 1) * 20)

	conditions := make([]bson.M, 0)

	conditions = append(conditions, bson.M{"$match": bson.M{"userId": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "tweet",
			"localfield":   "userrelationid",
			"foreignField": "userId",
			"as":           "tweet",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$tweet"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.date": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cur, err := col.Aggregate(ctx, conditions)
	var result []models.ReturnFollowTweets
	err = cur.All(ctx, &result)
	if err != nil {
		return result, false
	}
	return result, true
}
