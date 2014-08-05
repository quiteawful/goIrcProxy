package main

import (
	"errors"
	"strings"
	"time"
)

type Message struct {
	Timestamp time.Time
	User      string
	Content   string
}

var MessageLog []*Message

func msgWeb(user, content string) (bool, error) {

	user = strings.Trim(user, " ")

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

	con.Privmsg("#g0", "[web: "+user+"] "+content)
	return msgIrc(user, content)
}

func msgIrc(user, content string) (bool, error) {
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
