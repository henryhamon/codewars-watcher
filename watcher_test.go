package main

import (
	"math/rand"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var usernames = []string{"leometzger", "henryhammon", "ALNeneve"}

func TestAddUser(t *testing.T) {
	w := watcher
	n := rand.Intn(len(usernames))

	for _, u := range usernames {
		w.AddUser(u)
	}

	t.Log("\t\tchecking the AddUser method from watcher")

	t.Logf("\t\tchecking if length is correct after \"%d\" insertions \n", n)
	if len(watcher.Usernames) < n {
		t.Fatalf("\t\tshould length be \"%d\" gives \"%d\" insertions \"%s\" \n", len(watcher.Usernames), n, ballotX)
	}
	t.Log("\t\tshould be able to insert usernames", checkMark)

	t.Log("\t\tchecking the consistency of insertions")
	for i, u := range watcher.Usernames {
		if u != usernames[i] {
			t.Fatal("\t\tshould be able to maintain the consistency of insertions", ballotX)
		}
	}
	t.Log("\t\tshould be able to maintain the consistency of insertions", checkMark)
}

func TestAddUserError(t *testing.T) {
	t.Log("\t\tcheck if give an error when pass empty string")
	err := watcher.AddUser("")
	if err == nil {
		t.Fatal("should return an error when parameter is an empty string", ballotX)
	}
	t.Log("should return an error when parameter is an empty string", checkMark)
}
