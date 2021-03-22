package routers

import (
	"errors"
	"strings"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/* Email global variable */
var Email string

/* User global variable*/
var UserId string

/* Token Processing */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myKey := []byte("developerMaster_Twitter")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := db.CheckUserExists(claims.Email)
		if found {
			Email = claims.Email
			UserId = claims.ID.Hex()
		}
		return claims, found, UserId, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}
	return claims, false, string(""), err
}
