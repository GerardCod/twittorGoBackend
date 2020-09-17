package middlewares

import (
	"net/http"

	"github.com/GerardCod/twittorGoBackend/bd"
)

//CheckDB verifies if it already exists the user in signup
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.CheckConnection() == 0 {
			http.Error(w, "Connection lost.", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
