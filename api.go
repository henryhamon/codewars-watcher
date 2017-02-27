package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// API is a api to response with history of watcher users
type api struct {
	watcher *Watcher
	router  *mux.Router
}

func (a *api) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to codewars monitor api.")
}

// updateState - update state of all users
func (a *api) updateState(w http.ResponseWriter, r *http.Request) {
	watcher := GetWatcher()
	go watcher.UpdateUsers()
	w.WriteHeader(http.StatusNoContent)
}

// addUser - add a user to be monitored
func (a *api) addUser(w http.ResponseWriter, r *http.Request) {
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
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err = a.watcher.AddUser(user["username"]); err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(a.watcher); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
}

// Run - run api listener
func getapi() *api {
	var a api
	a.router = mux.NewRouter().StrictSlash(true)
	a.router.HandleFunc("/", a.index)
	a.router.HandleFunc("/update", a.updateState)
	a.watcher = GetWatcher()
	return &a
}
