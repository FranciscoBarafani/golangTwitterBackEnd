package routers

import (
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

/* New Relation Register */
func NewRelation(writter http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writter, "ID parameter missing", http.StatusBadRequest)
		return
	}

	var relation models.Relation
	relation.UserId = UserId
	relation.UserRelationId = ID

	status, err := db.NewRelation(relation)
	if err != nil {
		http.Error(writter, "Error while creating new relation", http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(writter, "Error while creating new relation", http.StatusBadRequest)
		return
	}
	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
