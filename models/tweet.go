package models

import "time"

/* Returning Tweet Model*/
type Tweet struct {
	Message string    `bson:"message" json:"message,omitempty"`
	Date    time.Time `bson:"date" json:"date,omitempty"`
}
