package main

import (
	"errors"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

// CompareState - compare with last state.
func CompareState(username string) ([]User, error) {
	var result []User
	session, err := mgo.Dial("")
	if err != nil {
		return nil, errors.New("")
	}
	defer session.Close()

	users := session.DB("codewars").C(UserCollection)
	err = users.
		Find(bson.M{"username": username}).
		Sort("-$natural").
		Limit(2).All(&result)
	return result, err
}
