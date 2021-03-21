package main

import (
	"log"

	db "github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/handlers"
)

func main() {
	if db.ConnectionCheck() == 0 {
		log.Fatal("No connection to database")
		return
	}
	handlers.Handlers()
}
