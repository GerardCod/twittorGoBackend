package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GerardCod/twittorGoBackend/bd"
	"github.com/GerardCod/twittorGoBackend/jwt"
	"github.com/GerardCod/twittorGoBackend/models"
)

//Login does the login with the DB
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "User or password invalid: "+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required.", 400)
		return
	}

	document, found := bd.Login(user.Email, user.Password)

	if !found {
		http.Error(w, "User does not exist or user and password invalid.", 400)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "There was an error with the token: "+err.Error(), 400)
		return
	}

	resp := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
