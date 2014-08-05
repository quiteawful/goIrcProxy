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
	user := r.FormValue("user")
	content := r.FormValue("content")

	if user != "" && content != "" { // is POST
		_, err := msgAdd(user, content)
		if err != nil {
			return
		}
	}

	fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>goIrcProxy</title></head><body>Log:<br>")
	for _, m := range MessageLog {
		fmt.Fprintf(w, "["+m.Timestamp.String()+"] "+m.User+": "+m.Content+"<br>")
	}

	fmt.Fprintf(w, "<form action='/' method='post'><input type='text' name='user'><input type='text' name='content'><input type='submit'></form></body></html>")
}
