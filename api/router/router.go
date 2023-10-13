package router

import (
	"handler/api/controller"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/profile/{username}", controller.ProfileHandler).Methods("GET")

	return router
}
