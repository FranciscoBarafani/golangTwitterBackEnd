package db

import (
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	"golang.org/x/crypto/bcrypt"
)

/* Checks if user exists to log in*/
func LoginTry(email string, password string) (models.User, bool) {
	/* Checks if user already exists*/
	user, userFound, _ := CheckUserExists(email)
	if !userFound {
		return user, false
	}
	/* If user is found we compare passwords*/
	passwordBytes := []byte(password)
	passwordDb := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDb, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
