package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/middlew"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Routes handling */
func Handlers() {
	router := mux.NewRouter()

	/*Routes*/
	router.HandleFunc("/register", middlew.DBcheck(routers.Register)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	/*CORS takes the router control */
	handler := cors.AllowAll().Handler(router)
	/*CORS takes the router control */
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
