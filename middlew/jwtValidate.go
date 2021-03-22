package middlew

import (
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/routers"
)

/* JWT Validation */
func JWTValidate(next http.HandlerFunc) http.HandlerFunc {
	return func(writter http.ResponseWriter, request *http.Request) {
		_, _, _, err := routers.ProcessToken(request.Header.Get("Authorization"))
		if err != nil {
			http.Error(writter, "Token error"+err.Error(), http.StatusBadRequest)
			return
		}
		/*IF the token is valid it sends the current request to the desired endpoint*/
		next.ServeHTTP(writter, request)
	}
}
