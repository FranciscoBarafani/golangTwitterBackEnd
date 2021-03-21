package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

/*Registers new User */
func Register(writter http.ResponseWriter, req *http.Request) {
	var user models.User
	/*Decodes request body string to JSON, and saves into err if there is an error*/
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		http.Error(writter, "Data error: "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(writter, "User email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(writter, "Password must have at least 6 characters", 400)
		return
	}

	_, found, _ := db.CheckUserExists(user.Email)
	if found == true {
		http.Error(writter, "An user with this email already exists", 400)
		return
	}

	_, status, err := db.InsertData(user)
	if err != nil {
		http.Error(writter, "Error when creating user: "+err.Error(), 400)
	}

	if status == false {
		http.Error(writter, "Error when creating user", 400)
	}

	writter.WriteHeader(http.StatusCreated)
}
