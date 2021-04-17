package db

import (
	"context"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

func DeleteRelation(relation models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Twitter")
	col := db.Collection("Relations")

	_, err := col.DeleteOne(ctx, relation)
	if err != nil {
		return false, err
	}
	return true, nil
}
