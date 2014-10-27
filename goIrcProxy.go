package main

var ctxIrc *Irc
var ctxLog *Log
var ctxUser map[string]string

func main() {
	ctxUser = make(map[string]string)
	ctxLog = new(Log)

	ctxIrc = new(Irc)

	ctxUser["vollgeheim"] = "soda"
	ctxUser["docschwammerl"] = "doclol"
	ctxIrc.Channels = append(ctxIrc.Channels, "#g0")
	ctxIrc.Network = "tardis.nerdlife.de"
	ctxIrc.Port = 6697

	go startWebServer(ctxUser)
	ctxIrc.Run()
}
