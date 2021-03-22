package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/jwt"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
)

/* User Login */
func Login(writter http.ResponseWriter, request *http.Request) {
	writter.Header().Add("content-type", "application/json")
	var user models.User

	/* Decodes user data from request body and inserts it into user variable */
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writter, "User or password are incorrect", 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(writter, "User email is required", 400)
		return
	}
	document, exists := db.LoginTry(user.Email, user.Password)
	if !exists {
		http.Error(writter, "User or password are incorrect", 400)
		return
	}

	/* JSON Web Token Generation */
	jwtKey, err := jwt.JWTGenerate(document)
	if err != nil {
		http.Error(writter, "An error ocurred when trying to generate token", 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	/* Writting response into header */
	writter.Header().Set("Content-Type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode(resp)

	/* Cookie Creation */
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(writter, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
