package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to codewars monitor api.")
}

// updateState - update state of all users
func updateState(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	go watcher.UpdateUsers()
	w.WriteHeader(http.StatusOK)
}

// addUser - add a user to be monitored
func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := make(map[string]string)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		return
	}
	if err = watcher.AddUser(user["username"]); err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(watcher.Usernames); err != nil {
		log.Println(err)
		return
	}
}

// removeUser - removes an user from watching user list
func removeUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	user := make(map[string]string)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(422)
		log.Println(err)
		return
	}

	if err = json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		return
	}

	if watcher.RemoveUser(user["username"]) {
		if err = json.NewEncoder(w).Encode(watcher.Usernames); err != nil {
			log.Println(err)
			return
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

// last retrieves last n registers from all users
func last(w http.ResponseWriter, r *http.Request) {
	var err error
	results := make([][]UserState, len(watcher.Usernames))
	vars := mux.Vars(r)
	n, _ := strconv.Atoi(vars["limit"])

	for i, u := range watcher.Usernames {
		results[i], err = watcher.Datastore.RegistersByLimit(u, n)
		if err != nil {
			log.Println(err)
		}
	}
	if err := json.NewEncoder(w).Encode(results); err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	if err := json.NewEncoder(w).Encode(watcher.Usernames); err != nil {
		log.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// RunAPI run api listener
func RunAPI(watcher Watcher) error {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/update", updateState)
	router.HandleFunc("/users", users)
	router.HandleFunc("/last/{limit:[0-9]+}", last)

	router.HandleFunc("/add", addUser)
	return http.ListenAndServe(":9090", router)
}
