package main

import "os/exec"

// Notifier - interface for OS to notify user
type Notifier interface {
	notify(message string)
}

// LinuxNotificator - linux treatment for notify users
type LinuxNotificator struct {
}

func (ln *LinuxNotificator) notify(message string) error {
	return exec.Command("notify-send", "codewars", message).Run()
}
