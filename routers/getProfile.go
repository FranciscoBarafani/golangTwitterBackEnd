package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func GetProfile(writter http.ResponseWriter, request *http.Request) {
	/*Gets ID from URL*/
	id := request.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(writter, "ID parameter is null", http.StatusBadRequest)
		return
	}
	profile, err := db.GetProfile(id)
	if err != nil {
		http.Error(writter, "An error ocurred while searching profile", http.StatusBadRequest)
		return
	}

	/*Writes Response Header*/
	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode(profile)
}
