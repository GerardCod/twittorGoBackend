package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/GerardCod/twittorGoBackend/middlewares"
	"github.com/GerardCod/twittorGoBackend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers contains the list of endpoints for the API
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/signup", middlewares.CheckDB(routers.SignUp)).Methods("POST")
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/showprofile", middlewares.CheckDB(middlewares.ValidateJWT(routers.ShowProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
