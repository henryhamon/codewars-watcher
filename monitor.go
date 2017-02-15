package main

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

// Monitor - monitor of friends.
type Monitor struct {
	usernames []string
	session   *mgo.Session
	time      time.Timer
}

// NewMonitor - monitor constructor
func NewMonitor(s *mgo.Session, unames []string) *Monitor {
	return &Monitor{session: s, usernames: unames}
}

func (m *Monitor) datastore() *UserStateDB {
	return &UserStateDB{m.session.Copy()}
}

// UserState - store a user and when he was in that state
type UserState struct {
	User User
	Date time.Time
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
