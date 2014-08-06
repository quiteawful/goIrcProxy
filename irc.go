package main

import (
	"crypto/tls"
	"fmt"
	"github.com/thoj/go-ircevent"
	"strconv"
)

type Irc struct {
	Con      *irc.Connection
	Network  string
	Port     int
	Channels []string
}

func (i *Irc) Run() {
	i.Con = irc.IRC("Datenkrake", "Datenkrake")
	i.Con.VerboseCallbackHandler = true
	i.Con.UseTLS = true
	i.Con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	i.Con.Connect(i.Network + ":" + strconv.Itoa(i.Port))

	i.Con.AddCallback("001", func(e *irc.Event) {
		i.Con.Join(i.Channels[0])
	})
	i.Con.AddCallback("PRIVMSG", parseIrcMsg)
	i.Con.AddCallback("CTCP_ACTION", parseIrcMsg)

	i.Con.Loop()
}

func (i *Irc) WriteToChannel(user, content string) {
	i.Con.Privmsg(i.Channels[0], fmt.Sprintf("[web: %s] %s", user, content))
}

func parseIrcMsg(e *irc.Event) {
	var user, content string
	if e.Code == "CTCP_ACTION" {
		user = "/me"
		content = e.Nick + " " + e.Arguments[1]
	}
	if e.Code == "PRIVMSG" {
		user = e.Nick
		content = e.Arguments[1]
	}
	ctxLog.AddIrcLog(user, content)
}
