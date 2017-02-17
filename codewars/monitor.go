package codewars

import (
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
)

var mon *Monitor

// Monitor - monitor of friends.
type Monitor struct {
	usernames []string
	session   *mgo.Session
	time      time.Timer
}

// GetMonitor - monitor constructor
func GetMonitor() (*Monitor, error) {
	if mon != nil {
		return mon, nil
	}
	s, err := mgo.Dial("")
	if err != nil {
		return nil, errors.New("error getting database session")
	}
	mon := &Monitor{session: s}
	return mon, err
}

// AddUser - add a user to be monitored
func (m *Monitor) AddUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	m.usernames = append(m.usernames, username)
	return nil
}

// UpdateUsers - update user state.
func (m *Monitor) UpdateUsers() error {
	//for _, username := range m.usernames {
	username := "henryhamon"
	user, err := GetUser(username)
	if err != nil {
		return err
	}
	userstate := UserState{User: user, Date: time.Now()}
	m.datastore().SaveUserState(userstate)
	//}
	return nil
}

func (m *Monitor) datastore() *UserStateDB {
	return &UserStateDB{m.session.Copy()}
}
