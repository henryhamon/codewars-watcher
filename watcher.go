package main

import (
	"errors"
	"time"
)

var w *Watcher

// UserState state of user when saved
type UserState struct {
	Time time.Time
	User User
}

// Watcher - Watcher of friends.
type Watcher struct {
	usernames []string
	datastore DataStore
	time      *time.Ticker
	client    CodewarsAPI
}

// GetWatcher - Watcher constructor
func GetWatcher() *Watcher {
	if w == nil {
		w = &Watcher{}
	}
	return w
}

// Run keep ticking and updating users
func (w *Watcher) Run(interrupt chan bool) {
	for {
		select {
		case <-w.time.C:
			w.UpdateUsers()
		case <-interrupt:
			return
		}
	}
}

// AddUser - add a user to be Watchered
func (w *Watcher) AddUser(username string) error {
	if username == "" {
		return errors.New("username cannot be empty")
	}
	w.usernames = append(w.usernames, username)
	return nil
}

// UpdateUsers - update user state.
func (w *Watcher) UpdateUsers() error {
	for _, username := range w.usernames {
		user, err := w.client.GetUser(username)
		if err != nil {
			return err
		}
		userstate := UserState{time.Now(), *user}
		err = w.datastore.Save(userstate)
		if err != nil {
			return err
		}
	}
	return nil
}
