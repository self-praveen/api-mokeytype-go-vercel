package controller

import (
	"fmt"
	"io/ioutil"
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
