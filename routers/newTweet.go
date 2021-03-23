package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

func NewTweet(writter http.ResponseWriter, request *http.Request) {
	var message models.Tweet

	err := json.NewDecoder(request.Body).Decode(&message)

	if err != nil {
		http.Error(writter, "Error while creating tweet "+err.Error(), 400)
		return
	}

	data := models.InsertTweet{
		UserID:  UserId,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.NewTweet(data)
	if err != nil {
		http.Error(writter, "Error while creating tweet "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(writter, "Error while creating tweet", 400)
		return
	}
	writter.WriteHeader(http.StatusCreated)
}
