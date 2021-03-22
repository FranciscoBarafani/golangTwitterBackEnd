package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

func EditProfile(writter http.ResponseWriter, request *http.Request) {
	var user models.User

	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writter, "Wrong data"+err.Error(), 400)
		return
	}

	status, err := db.EditProfile(user, UserId)
	if err != nil {
		http.Error(writter, "There has been an error while updating profile"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(writter, "Profile has not being updated"+err.Error(), 400)
		return
	}
	writter.WriteHeader(http.StatusCreated)
}
