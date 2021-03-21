package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = ConnectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://FranciscoBarafani:Barafan140407542@twitter-golang.rz74r.mongodb.net/test?authSource=admin&replicaSet=atlas-5mua6n-shard-0&readPreference=primary&appname=MongoDB%20Compass&ssl=true")

/*Database connection*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Database connection succesful")
	return client
}

/*Checks connection to database */
func ConnectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
