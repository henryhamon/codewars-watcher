package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// UserStateCol - collection name for saving users
const UserStateCol string = "userstate"

// DataBase - database name
const DataBase string = "codewars"

// DataStore interface to store data saved from api
type DataStore interface {
	Save(us UserState) error
	RegistersByDate(username string, time time.Time) ([]UserState, error)
	RegistersByLimit(username string, limit int) ([]UserState, error)
}

// MongoStore - struct to databasing the application
type MongoStore struct {
	session *mgo.Session
}

// UserstateCol - gives a collection of userstate.
func (ds *MongoStore) userstateCol() *mgo.Collection {
	return ds.session.DB(DataBase).C(UserStateCol)
}

// Save - saves a user state into database
func (ds *MongoStore) Save(us UserState) error {
	return ds.userstateCol().Insert(&us)
}

// RegistersByLimit - return last n states of user.
func (ds *MongoStore) RegistersByLimit(username string, n int) ([]UserState, error) {
	var uss []UserState
	err := ds.userstateCol().Find(bson.M{"username": username}).Sort("-$natural").Limit(n).All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by username")
	}
	return uss, err
}

// RegistersByDate - get last dates user states
func (ds *MongoStore) RegistersByDate(username string, t time.Time) ([]UserState, error) {
	var uss []UserState
	err := ds.userstateCol().Find(bson.M{"username": username, "date": t}).Sort("-$natural").All(&uss)

	if err != nil {
		return nil, errors.New("error finding users by date")
	}
	return uss, err
}

// FileStore an struct to save states into a file
type FileStore struct {
	path   string
	states []UserState
}

// NewFileStore file store constructor
func NewFileStore() *FileStore {
	dir, err := os.Getwd()
	path := dir + "/tmp"
	if err != nil {
		log.Fatal(err)
	}
	if !exists(path) {
		err = os.Mkdir(dir+"/tmp", os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	return &FileStore{path: path}
}

// Save implementation to save into a file
func (ds *FileStore) Save(us UserState) error {
	completePath := ds.path + "/" + us.User.Username + ".txt"
	data, err := json.Marshal(us)
	if err != nil {
		return err
	}
	err = saveFile(completePath, data)
	return err
}

// RegistersByDate retrieves last n register registered from time t
func (ds *FileStore) RegistersByDate(username string, t time.Time) ([]UserState, error) {
	return nil, nil
}

// RegistersByLimit retrieves las n registers n
func (ds *FileStore) RegistersByLimit(username string, limit int) ([]UserState, error) {
	return nil, nil
}

func saveFile(name string, data []byte) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
