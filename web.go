package main

import (
	"fmt"
	"net/http"
)

var HtmlTemplate string = "<!DOCTYPE html><html><head><title>goIrcProxy</title></head><body>{{Log}}<form action='/' method='post'><input type='text' name='user'><input type='text' name='content'><input type='submit'></form></body></html>"

func startWebServer() {
	http.HandleFunc("/", HttpIndex)
	http.ListenAndServe(":8080", nil)
}

// Serves the main index page, pretty much only the messagelog
func HttpIndex(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")

	if content != "" { // is POST
		ctxLog.AddWebLog("doclol", content)
	}

	fmt.Fprintf(w, "<!DOCTYPE html><html><head><meta http-equiv='refresh' content='5'><title>goIrcProxy</title></head><body>")
	for _, m := range ctxLog.MessageLog {
		fmt.Fprintf(w, fmt.Sprintf("[%s] %s: %s<br>", m.Timestamp.String(), m.User, m.Content))
	}

	fmt.Fprintf(w, "<form action='/' method='post'>Nachricht: <input type='text' name='content'><input type='submit'></form></body></html>")
}
