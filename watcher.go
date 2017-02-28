package main

import (
	"errors"
	"time"
)

// UserState state of user when saved
type UserState struct {
	Time time.Time
	User User
}

// Watcher - Watcher of friends.
type Watcher struct {
	Usernames []string     `json:"usernames"`
	Datastore DataStore    `json:"mongo"`
	Time      *time.Ticker `json:"ticker"`
	client    CodewarsAPI
}

// Run keep ticking and updating users
func (w *Watcher) Run(stop chan bool) {
	for {
		select {
		case <-w.Time.C:
			w.UpdateUsers()
		case <-stop:
			return
		}
	}
}

// AddUser - add a user to be Watchered
func (w *Watcher) AddUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	w.Usernames = append(w.Usernames, username)
	return nil
}

// RemoveUser - remove an user from list of users
func (w *Watcher) RemoveUser(username string) bool {
	for i, u := range w.Usernames {
		if u == username {
			w.Usernames = append(w.Usernames[:i], w.Usernames[i+1:]...)
			return true
		}
	}
	return false
}

// UpdateUsers - update user state.
func (w *Watcher) UpdateUsers() error {
	for _, username := range w.Usernames {
		user, err := w.client.GetUser(username)
		if err != nil {
			return err
		}
		userstate := UserState{time.Now(), *user}
		err = w.Datastore.Save(userstate)
		if err != nil {
			return err
		}
	}
	return nil
}
