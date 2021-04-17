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
	router.HandleFunc("/login", middlew.DBcheck(routers.Login)).Methods("POST")
	router.HandleFunc("/getProfile", middlew.DBcheck(middlew.JWTValidate(routers.GetProfile))).Methods("GET")
	router.HandleFunc("/editProfile", middlew.DBcheck(middlew.JWTValidate(routers.EditProfile))).Methods("PUT")
	router.HandleFunc("/newTweet", middlew.DBcheck(middlew.JWTValidate(routers.NewTweet))).Methods("POST")
	router.HandleFunc("/getTweet", middlew.DBcheck(middlew.JWTValidate(routers.GetTweet))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlew.DBcheck(middlew.JWTValidate(routers.DeleteTweet))).Methods("DELETE")
	router.HandleFunc("/newRelation", middlew.DBcheck(middlew.JWTValidate(routers.NewRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlew.DBcheck(middlew.JWTValidate(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/getRelation", middlew.DBcheck(middlew.JWTValidate(routers.GetRelation))).Methods("GET")

	router.HandleFunc("/getAvatar", middlew.DBcheck(middlew.JWTValidate(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getBanner", middlew.DBcheck(middlew.JWTValidate(routers.GetBanner))).Methods("GET")
	router.HandleFunc("/uploadAvatar", middlew.DBcheck(routers.UploadAvatar)).Methods("POST")
	router.HandleFunc("/uploadBanner", middlew.DBcheck(routers.UploadAvatar)).Methods("POST")
	router.HandleFunc("/listUsers", middlew.DBcheck(middlew.JWTValidate(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	/*CORS takes the router control */
	handler := cors.AllowAll().Handler(router)
	/*CORS takes the router control */
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
