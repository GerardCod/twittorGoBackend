package routers

import (
	"encoding/json"
	"net/http"

	"github.com/GerardCod/twittorGoBackend/bd"
)

//ShowProfile retrieves data about the user given an id.
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "Id is required.", http.StatusBadRequest)
		return
	}

	profile, err := bd.ShowProfile(id)

	if err != nil {
		http.Error(w, "There was an error trying to search the profile", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}
