package main

import (
	"github.com/0xAX/notificator"
)

func notify(message string) {
	var notify *notificator.Notificator
	notify = notificator.New(notificator.Options{
		DefaultIcon: "",
		AppName:     "Codewars",
	})
	notify.Push("title", message, "", notificator.UR_NORMAL)
}
