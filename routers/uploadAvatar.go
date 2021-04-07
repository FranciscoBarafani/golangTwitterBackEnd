package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

/* Upload Avatar Route */
func UploadAvatar(writter http.ResponseWriter, request *http.Request) {
	/* Gets the file from the request */
	file, handler, _ := request.FormFile("avatar")
	/* We split the filename from the extension */
	var extension = strings.Split(handler.Filename, ".")[1]
	/* File name set up */
	var fileName string = "uploads/avatars" + UserId + "." + extension

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(writter, "Error while uploading image"+err.Error(), http.StatusBadRequest)
		return
	}
	/* Saves file in storage */
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(writter, "Error while saving image"+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool
	user.Avatar = UserId + "." + extension
	status, err = db.EditProfile(user, UserId)
	if err != nil || !status {
		http.Error(writter, "Error while saving image in user"+err.Error(), http.StatusBadRequest)
		return
	}

	writter.Header().Set("content-type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
