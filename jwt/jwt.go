package jwt

import (
	"time"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*JWT Generation*/
func JWTGenerate(user models.User) (string, error) {
	myKey := []byte("developerMaster_Twitter")
	payload := jwt.MapClaims{
		"email":    user.Email,
		"name":     user.Name,
		"lastName": user.Lastname,
		"bio":      user.Bio,
		"webSite":  user.WebSite,
		"_id":      user.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
