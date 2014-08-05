package main

import (
	"errors"
	"time"
)

type Message struct {
	Timestamp time.Time
	User      string
	Content   string
}

var MessageLog []*Message

func msgAdd(user, content string) (bool, error) {
	if user == "" {
		return false, errors.New("Username is empty.")
	}
	if len(user) > 20 {
		return false, errors.New("Username is too long. Max: 20")
	}
	if content == "" {
		return false, errors.New("Content is empty.")
	}
	if len(content) > 100 {
		return false, errors.New("Content is too long. Max: 100")
	}
	if !isWhitelisted(user) {
		return false, errors.New("Sorry, only for VIPs.")
	}

	checkLimits()
	m := &Message{Timestamp: time.Now(), User: user, Content: content}

	MessageLog = append(MessageLog, m)

	return true, nil
}

func checkLimits() {
	if len(MessageLog) > 100 {
		MessageLog = MessageLog[1:]
	}
}
