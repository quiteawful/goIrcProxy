package main

import (
	"crypto/tls"
	"fmt"
	"github.com/thoj/go-ircevent"
)

var con *irc.Connection

func main() {
	fmt.Println("Start Proxy")
	//MessageLog = append(MessageLog, &Message{Timestamp: time.Now(), User: "marduk", Content: "hi"})
	go startWebServer()

	con = irc.IRC("Datenkrake", "Datenkrake")

	con.VerboseCallbackHandler = true
	con.UseTLS = true
	con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := con.Connect("tardis.nerdlife.de:6697")

	if err != nil {
		fmt.Println("error while connecting.")
	}

	// Join Channel
	con.AddCallback("001", func(e *irc.Event) {
		con.Join("#g0")
	})

	con.AddCallback("PRIVMSG", func(e *irc.Event) {
		msgIrc(e.Nick, e.Arguments[1])
	})

	con.Loop()
	ch := make(chan bool)
	<-ch
	fmt.Println("heeey")
}
