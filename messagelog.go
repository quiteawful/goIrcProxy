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
	content = html.EscapeString(content)
	l.MessageLog = append(l.MessageLog, NewMessage(user, content))
}

func (l *Log) checkLimit() {
	// backlog of 50
	if len(l.MessageLog) > 50 {
		l.MessageLog = l.MessageLog[len(l.MessageLog)-50:]
	}
}
