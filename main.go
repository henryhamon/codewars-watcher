package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// USERNAME - user to monitor
const USERNAME string = "henryhamon"

// SERVICE - service url for monitor a friend
const SERVICE string = "https://www.codewars.com/api/v1/users/"

func main() {

	var user User

	resp, err := http.Get(SERVICE + USERNAME)
	if err != nil {
		log.Fatal("Get service ", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal("Decoder ", err)
	}

	if err := SaveState(user); err != nil {
		log.Fatal(err)
	}
}
