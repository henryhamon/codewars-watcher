package main

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
)

// UserCollection - collection name for saving users
const UserCollection string = "users"

// DataBase - database name
const DataBase string = "codewars"

// SaveState - saves a state of a user into database
func SaveState(u User) error {
	session, err := mgo.Dial("")
	if err != nil {
		return errors.New("error getting a database session")
	}
	defer session.Close()

	users := session.DB("codewars").C(UserCollection)
	if err := users.Insert(&u); err != nil {
		return errors.New("error inserting user into database")
	}
	return nil
}
