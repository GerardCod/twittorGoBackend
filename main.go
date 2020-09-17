package main

import (
	"log"

	"github.com/GerardCod/twittorGoBackend/bd"
	"github.com/GerardCod/twittorGoBackend/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("No connection.")
		return
	}
	handlers.Handlers()
}
