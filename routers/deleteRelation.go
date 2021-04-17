package routers

import (
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

func DeleteRelation(writter http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	var relation models.Relation
	relation.UserId = UserId
	relation.UserRelationId = ID

	status, err := db.DeleteRelation(relation)
	if err != nil {
		http.Error(writter, "Error while deleting relation", http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(writter, "Error while deleting relation", http.StatusBadRequest)
		return
	}
	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
