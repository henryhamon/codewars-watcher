package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome to codewars monitor api.")
}

// updateState - update state of all users
func updateState(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	go watcher.UpdateUsers()
	w.WriteHeader(http.StatusOK)
}

// addUser - add a user to be monitored
func addUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	user := make(map[string]string)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Println(err)
		w.WriteHeader(422)
		return
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
func removeUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

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

func updateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var userwh UserWebhook

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		w.WriteHeader(422)
		log.Println(err)
		return
	}
	if err = json.Unmarshal(body, &userwh); err != nil {
		w.WriteHeader(422)
		log.Println(err)
		return
	}
	_ = watcher.UpdateUser(userwh.User.Username)
	w.WriteHeader(http.StatusOK)
}

// last retrieves last n registers from all users
func last(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var err error
	results := make([][]UserState, len(watcher.Usernames))
	n, _ := strconv.Atoi(ps.ByName("limit"))

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

func users(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
	router := httprouter.New()

	router.GET("/api/v1/", index)
	router.GET("/api/v1/update", updateState)
	router.GET("/api/v1/users", users)
	router.GET("/api/v1/last/{limit:[0-9]+}", last)

	router.POST("/api/v1/add", addUser)
	router.POST("/api/v1/update/user", updateUser)

	return http.ListenAndServe(":9090", router)
}
