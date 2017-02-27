package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var watcher *Watcher

func main() {
	watcher = GetWatcher()
	interrupt := make(chan bool)
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	usernames := []string{"leometzger", "henryhamon"}

	watcher.datastore = GetDataStore(session)
	watcher.time = time.NewTicker(10 * time.Second)
	watcher.usernames = usernames
	go watcher.Run(interrupt)

	fmt.Println("listening on the port 9090 localhost")
	var server api
	err = http.ListenAndServe(":9090", server.router)
	if err != nil {
		interrupt <- true
		log.Fatal(err)
	}
}

// GetDataStore gives the datastore correct
func GetDataStore(session *mgo.Session) DataStore {
	return MongoStore{session}
}
