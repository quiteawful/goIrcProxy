package main

import (
	"html"
	"time"
)

type Message struct {
	Timestamp time.Time
	User      string
	Content   string
}

func NewMessage(user, content string) *Message {
	return &Message{Timestamp: time.Now(), User: user, Content: content}
}

type Log struct {
	MessageLog []*Message
}

// add a new Log to
func (l *Log) AddWebLog(user, content string) {
	ctxIrc.WriteToChannel(user, content)
	l.AddIrcLog(user, content)
}

func (l *Log) AddIrcLog(user, content string) {
	l.checkLimit()
	// es wird der Input von Irc-Usern und von Web-Usern gefiltert!
	// wahlweise kann man wohl auch template.HTMLEscapeString() nehmen
	content = html.EscapeString(content)
	user = html.EscapeString(user)
	l.MessageLog = append(l.MessageLog, NewMessage(user, content))
}

func (l *Log) checkLimit() {
	var numBacklog int = 50
	// backlog of 50
	if len(l.MessageLog) > numBacklog {
		l.MessageLog = l.MessageLog[len(l.MessageLog)-numBacklog:]
	}
}
