package main

import (
	"handler/api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/profile/{username}", controller.ProfileHandler).Methods("GET")

	return router
}

func main() {
	r := Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}
