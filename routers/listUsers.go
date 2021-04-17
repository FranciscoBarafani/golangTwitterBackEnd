package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func ListUsers(writter http.ResponseWriter, req *http.Request) {

	kind := req.URL.Query().Get("kind")
	page := req.URL.Query().Get("page")
	search := req.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(writter, "You must send as parameter a page number greater than 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)
	result, status := db.GetAllUsers(UserId, pag, search, kind)
	if !status {
		http.Error(writter, "Error while reading users", http.StatusBadRequest)
		return
	}
	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode(result)
}
