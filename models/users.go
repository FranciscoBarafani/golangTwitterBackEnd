package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Database model creation, first we declare the attribute name, then type.
Aftewards we declare the bson data, how is it going to be sent to mongodb. Entry Data
Then the json is to declare how is it going to be returned from the db. Exit Data
*/
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Lastname string             `bson:"lastname" json:"lastname,omitempty"`
	BornDate time.Time          `bson:"bornDate" json:"bornDate,omitempty"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password,omitempty"`
	Avatar   string             `bson:"avatar" json:"avatar,omitempty"`
	Banner   string             `bson:"banner" json:"banner,omitempty"`
	Bio      string             `bson:"bio" json:"bio,omitempty"`
	WebSite  string             `bson:"webSite" json:"webSite,omitempty"`
}
