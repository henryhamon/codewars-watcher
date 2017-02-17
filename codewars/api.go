package codewars

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to codewars monitor api.")
}

// updateState - update state of all users
func updateState(w http.ResponseWriter, r *http.Request) {
	monitor, err := GetMonitor()
	if err != nil {
	}
	log.Fatal(err)
	go monitor.UpdateUsers()
	w.WriteHeader(http.StatusNoContent)
}

// addUser - add a user to be monitored
// NOT RETURNING.
func addUser(w http.ResponseWriter, r *http.Request) {
	monitor, err := GetMonitor()
	if err != nil {
		log.Fatal(err)
	}
	user := make(map[string]string)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(422)
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = monitor.AddUser(user["username"]); err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(monitor); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Run - run api listener
func Run() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/add", addUser)
	router.HandleFunc("/update", updateState)

	log.Fatal(http.ListenAndServe(":8989", router))
}
