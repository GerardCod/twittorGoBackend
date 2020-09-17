package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GerardCod/twittorGoBackend/bd"
	"github.com/GerardCod/twittorGoBackend/models"
)

//SignUp create a new User in DB
func SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Error in data sent:"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "Email is required.", 400)
		return
	}

	if len(user.Password) < 8 {
		http.Error(w, "Password must have 8 characters minimun.", 400)
		return
	}

	_, found, _ := bd.UserAlreadyExists(user.Email)

	if found {
		http.Error(w, "There is a count with this email.", 400)
		return
	}

	_, status, err := bd.SignUp(user)

	if err != nil {
		http.Error(w, "There was an error: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Signup failed", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
