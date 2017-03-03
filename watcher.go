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

// UpdateUser updates an specific user state
func (w *Watcher) UpdateUser(usernameOrID string) error {
	user, err := w.client.GetUser(usernameOrID)
	if err != nil {
		return err
	}
	userstate := UserState{time.Now(), *user}
	changed, err := w.UserChanged(*user)
	if err != nil {
		return err
	}
	if !changed {
		return nil
	}
	err = w.Datastore.Save(userstate)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUsers - update user state.
func (w *Watcher) UpdateUsers() error {
	for _, username := range w.Usernames {
		err := w.UpdateUser(username)
		if err != nil {
			return err
		}
	}
	return nil
}

// UserChanged was made for check if an user change since last time your state
// was saved
func (w *Watcher) UserChanged(user User) (bool, error) {
	users, err := w.Datastore.RegistersByLimit(user.Username, 1)
	if err != nil {
		return true, errors.New("error searching user")
	}
	if len(users) < 1 {
		return true, nil
	}
	honor, leaderboard := CompareUsers(users[0].User, user)
	return !(honor == 0 && leaderboard == 0), nil
}

// CompareUsers returns a difference of honor and leaderboard between them
// based on scores of user 1
func CompareUsers(u1, u2 User) (honor, leaderboard int) {
	return u1.Honor - u2.Honor, u1.LeaderboardPosition - u2.LeaderboardPosition
}
