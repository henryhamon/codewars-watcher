package main

import mgo "gopkg.in/mgo.v2"

// USERNAME - user to monitor
const USERNAME string = "henryhamon"

// SERVICE - service url for monitor a friend
const SERVICE string = "https://www.codewars.com/api/v1/users/"

var (
	username string
	friends  []string
	user     User
)

func main() {

	usernames := make([]string, 5)
	usernames = append(usernames, username)

	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	monitor := NewMonitor(session, usernames)
	err = monitor.UpdateUsers()
	if err != nil {
		panic(err)
	}
}
