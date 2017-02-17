package codewars

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

// UserStateDB - struct to databasing the application
type UserStateDB struct {
	session *mgo.Session
}

// UserState - store a user and when he was in that state
type UserState struct {
	User User
	Date time.Time
}

// UserstateCol - gives a collection of userstate.
func (ds *UserStateDB) UserstateCol() *mgo.Collection {
	return ds.session.DB(DataBase).C(UserStateCol)
}

// SaveUserState - saves a user state into database
func (ds *UserStateDB) SaveUserState(us UserState) error {
	return ds.UserstateCol().Insert(&us)
}

// GetByLimit - return last n states of user.
func (ds *UserStateDB) GetByLimit(username string, n int) ([]UserState, error) {
	var uss []UserState
	err := ds.UserstateCol().Find(bson.M{"username": username}).Sort("-$natural").Limit(n).All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by username")
	}
	return uss, err
}

// GetByDate - get last dates user states
func (ds *UserStateDB) GetByDate(username string, t time.Time) ([]UserState, error) {
	var uss []UserState
	err := ds.UserstateCol().Find(bson.M{"username": username, "date": t}).Sort("-$natural").All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by date")
	}
	return uss, err
}
