package main

import (
	"handler/api/router"
	"log"
	"net/http"
)

func main() {
	r := router.Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}
