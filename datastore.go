package main

import (
	"errors"
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserStateCol - collection name for saving users
const UserStateCol string = "userstate"

// DataBase - database name
const DataBase string = "codewars"

// UserStateDB - struct to databasing the application
type UserStateDB struct {
	session *mgo.Session
}

// UserstateCol - gives a collection of userstate.
func (ds *UserStateDB) UserstateCol() *mgo.Collection {
	return ds.session.DB(DataBase).C(UserStateCol)
}

// SaveUserState - saves a user state into database
func (ds *UserStateDB) SaveUserState(us UserState) error {
	fmt.Println(us)
	err := ds.UserstateCol().Insert(&us)
	return err
}

// GetUserStates - return last n states of user.
func (ds *UserStateDB) GetUserStates(username string, n int) ([]UserState, error) {

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
func (ds *UserStateDB) GetUserStatesByTime(username string, t time.Time) ([]UserState, error) {

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
