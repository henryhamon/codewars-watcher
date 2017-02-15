package main

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserStateCol - collection name for saving users
const UserStateCol string = "userstate"

// DataBase - database name
const DataBase string = "codewars"

// DataStore - struct to databasing the application
type DataStore struct {
	session *mgo.Session
}

// UserstateCol - gives a collection of userstate.
func (ds *DataStore) UserstateCol() *mgo.Collection {
	return ds.session.DB(DataBase).C(UserStateCol)
}

// SaveUserState - saves a user state into database
func (ds *DataStore) SaveUserState(us UserState) error {
	err := ds.UserstateCol().Insert(&us)
	return err
}

// GetUserStates - return last n states of user.
func (ds *DataStore) GetUserStates(username string, n int) ([]UserState, error) {

	var uss []UserState

	col := ds.UserstateCol()
	err := col.Find(bson.M{"username": username}).
		Sort("-$natural").
		Limit(n).
		All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by username")
	}
	return uss, err
}

// GetUserStatesByTime - get last dates user states
func (ds *DataStore) GetUserStatesByTime(username string, t time.Time) ([]UserState, error) {

	var uss []UserState

	col := ds.UserstateCol()
	err := col.Find(bson.M{"username": username, "date": t}).
		Sort("-$natural").
		All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by username")
	}
	return uss, err
}
