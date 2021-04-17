package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func GetTweetsRelation(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send page param", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must send page parameter greater than 0", http.StatusBadRequest)
		return
	}

	response, correct := db.GetFollowTweets(UserId, page)
	if !correct {
		http.Error(w, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
