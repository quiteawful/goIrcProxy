package main

import (
	"fmt"
)

var ctxIrc *Irc
var ctxLog *Log

func main() {
	ctxLog = new(Log)

	ctxIrc = new(Irc)
	ctxIrc.Channels = append(ctxIrc.Channels, "#g0")
	ctxIrc.Network = "tardis.nerdlife.de"
	ctxIrc.Port = 6697

	go startWebServer()
	ctxIrc.Run()

	fmt.Println("done.")
}
