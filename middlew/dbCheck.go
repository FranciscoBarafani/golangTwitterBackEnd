package middlew

import (
	"net/http"

	"github.com/FranciscoBarafani/golangTwitterBackEnd/db"
)

/*Checks connection to database prior to execute query */
func DBcheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*If there is an error to connect it returns error 500 */
		if db.ConnectionCheck() == 0 {
			http.Error(w, "Database connection lost", 500)
			return
		}
		/*If it is success it takes the data to the next step */
		next.ServeHTTP(w, r)
	}
}
