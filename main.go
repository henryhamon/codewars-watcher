package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var watcher *Watcher
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	watcher = GetWatcher()
	interrupt := make(chan bool)
	session, err := mgo.Dial("")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	usernames := []string{"leometzger", "henryhamon", "ALNeneve"}

	watcher.datastore = GetDataStore(session)
	watcher.time = time.NewTicker(3600 * time.Second)
	watcher.usernames = usernames
	go watcher.Run(interrupt)

	fmt.Println("watching your friends and updating hourly")
	wg.Wait()
}

// GetDataStore gives the datastore correct
func GetDataStore(session *mgo.Session) DataStore {
	return MongoStore{session}
}
