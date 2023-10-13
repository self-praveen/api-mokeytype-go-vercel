package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	apiURL := fmt.Sprintf("https://api.monkeytype.com/users/%s/profile", username)

	response, err := http.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch data from Monkeytype API", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, "Failed to read response from Monkeytype API", http.StatusInternalServerError)
		return
	}

	if response.StatusCode != http.StatusOK {
		http.Error(w, "Failed to fetch data from Monkeytype API. User not found or other error.", response.StatusCode)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseBody)
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/profile/{username}", ProfileHandler).Methods("GET")
	return router
}

func main() {
	r := Router()

	err := http.ListenAndServe(":4068", r)
	if err != nil {
		log.Fatal(err)
	}
}
