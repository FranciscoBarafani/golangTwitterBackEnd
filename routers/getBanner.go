package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

func GetBanner(writter http.ResponseWriter, request *http.Request) {

	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(writter, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	profile, err := db.GetProfile(ID)
	if err != nil {
		http.Error(writter, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(writter, "Image not found", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(writter, OpenFile)
	if err != nil {
		http.Error(writter, "Error while getting image", http.StatusBadRequest)
	}
}
