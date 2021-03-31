package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func GetTweet(writter http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writter, "ID param not defined", http.StatusBadRequest)
		return
	}
	if len(request.URL.Query().Get("page")) < 1 {
		http.Error(writter, "Page number param not defined", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		http.Error(writter, "Page number format error", http.StatusBadRequest)
		return
	}
	pageNumber := int64(page)
	response, correct := db.GetTweet(ID, pageNumber)
	if !correct {
		http.Error(writter, "Error while reading tweets", http.StatusBadRequest)
		return
	}
	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode(response)
}
