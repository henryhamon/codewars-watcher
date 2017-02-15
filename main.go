package main

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
	/*
		session, err := mgo.Dial("")
		if err != nil {
			panic(err)
		}
		defer session.Close()
	*/
	ln := LinuxNotificator{}
	err := ln.notify("testano")
	if err != nil {
		panic(err)
	}
}
