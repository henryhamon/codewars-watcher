package main

import (
	"encoding/json"
	"errors"
	"net/http"
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

func (m *Monitor) dataStore() *DataStore {
	return &DataStore{m.session.Copy()}
}

// UserState - store a user and when he was in that state
type UserState struct {
	user User
	date time.Time
}

// UpdateUsers - update user state.
func (m *Monitor) UpdateUsers() error {
	for _, username := range m.usernames {
		user, err := m.getUser(username)
		if err != nil {
			return err
		}
		userstate := UserState{user, time.Now()}
		m.dataStore().SaveUserState(userstate)
	}
	return nil
}

func (m *Monitor) getUser(username string) (User, error) {
	var user User
	resp, err := http.Get(SERVICE + username)
	if err != nil {
		return user, errors.New("error getting a user")
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return user, errors.New("error decoding a user")
	}
	return user, err
}
