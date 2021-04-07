package routers

import (
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func DeleteTweet(writter http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writter, "ID must be sent", http.StatusBadRequest)
		return
	}

	err := db.DeleteTweet(ID, UserId)
	if err != nil {
		http.Error(writter, "There was an error while deleting tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
