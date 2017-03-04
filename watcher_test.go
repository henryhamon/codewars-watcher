package main

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var usernames = []string{"leometzger", "henryhammon", "ALNeneve"}
var userstates = [][]UserState{
	{UserState{
		Time: time.Now(),
		User: User{
			Username:            "leometzger",
			Name:                "Leonardo Metzger",
			Honor:               1525,
			Clan:                "BPlus",
			LeaderboardPosition: 1723,
		},
	},
	},
}

// mock implementation of Datastore,
// just returs fixed users
type mockDatastore struct {
	userstates [][]UserState
}

func newMockDatastore() *mockDatastore {
	return &mockDatastore{userstates}
}

func (ds *mockDatastore) Save(us UserState) error {
	return nil
}

func (ds *mockDatastore) RegistersByLimit(username string, limit int) ([]UserState, error) {
	return ds.userstates[limit-1], nil
}

func TestAddUser(t *testing.T) {
	for _, u := range usernames {
		watcher.AddUser(u)
	}

	assert.Equal(t, len(usernames), len(watcher.Usernames), "length users should be equal")

	for i, u := range watcher.Usernames {
		assert.Equal(t, usernames[i], u, "names should be equal")
	}

	err := watcher.AddUser("")
	assert.NotNil(t, err, "error should not be nil when empty string is sended")
}

func TestCompareUsers(t *testing.T) {
	u1, u2 := User{}, User{}
	honor, leaderboard := CompareUsers(u1, u2)

	assert.Equal(t, honor, 0, "they should be equal")
	assert.Equal(t, leaderboard, 0, "they should be equal")
}

func TestUserChanged(t *testing.T) {

	watcher.Datastore = newMockDatastore()
	user := userstates[0][0].User

	actual, err := watcher.UserChanged(user)
	if assert.Nil(t, err, "error should be nil") {
		assert.False(t, actual, "changed should be false")
	}

	user.Honor = user.Honor + rand.Intn(200)
	user.LeaderboardPosition = user.LeaderboardPosition - rand.Intn(1500)

	actual, err = watcher.UserChanged(user)
	if assert.Nil(t, err, "error should be nil") {
		assert.True(t, actual, "changed should be true")
	}
}
