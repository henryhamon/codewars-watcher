package main

import (
	"fmt"

	"github.com/leometzger/codewars-monitor/codewars"
)

func main() {
	/*
		session, err := mgo.Dial("")
		if err != nil {
			panic(err)
		}
		defer session.Close()
	*/
	fmt.Println("listening on port 8989")
	codewars.Run()
	fmt.Println("closing api.")
}
