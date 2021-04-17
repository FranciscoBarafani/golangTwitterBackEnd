package routers

import (
	"encoding/json"
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

func GetRelation(writter http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	var relation models.Relation
	relation.UserId = UserId
	relation.UserRelationId = ID

	var resp models.GetRelationResponse

	status, err := db.GetRelation(relation)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode(resp)
}
