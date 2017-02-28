package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
	"time"

	mgo "gopkg.in/mgo.v2"
)

const configfile = "config_file.json"

//  Package variables
var config configFile
var watcher Watcher
var session *mgo.Session

func init() {

	config = defaultConfigFile()

	if !exists(configfile) {
		return
	}

	data, err := ioutil.ReadFile(configfile)
	fmt.Println(string(data))
	if err != nil {
		return
	}

	if err = json.Unmarshal(data, &config); err != nil {
		log.Println("error unmarshaling config file", err)
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	dtstore := GetDataStore(session)
	timer := time.NewTicker(time.Duration(config.Ticker) * time.Hour)

	usernames := []string{"leometzger", "henryhamon", "ALNeneve"}

	watcher = Watcher{
		Datastore: dtstore,
		client:    CodewarsAPI{},
		Time:      timer,
		Usernames: usernames,
	}
	stop := make(chan bool)
	go watcher.Run(stop)

	fmt.Printf("watching your friends and updating each %d hours\n", config.Ticker)

	if config.API {
		fmt.Println("Running api on port 9090")
		err := RunAPI(watcher)
		if err != nil {
			log.Fatal(err)
			stop <- true
		}
	}
	wg.Wait()
}

// GetDataStore gives the datastore configured
func GetDataStore(session *mgo.Session) DataStore {

	if config.Datastore == "Mongo" {
		session, err := mgo.Dial(config.DatabaseURL)

		if err != nil {
			log.Fatal(err)
		}
		return MongoStore{session.Copy()}
	}
	return FileStore{}
}
