package main

import (
	"fmt"
	"net/http"
)

func startWebServer() {
	http.HandleFunc("/", HttpIndex)
	http.HandleFunc("/write", HttpWrite)
	http.ListenAndServe(":8080", nil)
}

// Serves the main index page, pretty much only the messagelog
func HttpIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<html><head></head><body>")
	if len(MessageLog) > 0 {
		for _, m := range MessageLog {
			fmt.Fprintf(w, "[%s] %s: %s<br>", m.Timestamp.String(), m.User, m.Content)
		}
	}
	fmt.Fprintf(w, FormTemplate+"</body></html>")
}

func HttpWrite(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	content := r.FormValue("content")

	if user == "" || content == "" {
		return
	}

	// write to irc
	_, err := msgAdd(user, content)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
