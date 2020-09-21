package middlewares

import "net/http"

//ValidateJWT is a middleware that verifies if a jwt is valid.
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token error: "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
