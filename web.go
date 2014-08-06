package main

import (
	"fmt"
	"net/http"
)

func startWebServer() {
	http.HandleFunc("/", HttpIndex)
	http.HandleFunc("/log", HttpLog)
	http.ListenAndServe(":8080", nil)
}

// Serves the main index page, pretty much only the messagelog
func HttpIndex(w http.ResponseWriter, r *http.Request) {
	content := r.FormValue("content")

	if content != "" { // is POST
		ctxLog.AddWebLog("doclol", content)
	}

	fmt.Fprintf(w, HtmlMain)
}

func HttpLog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, HtmlLogHead)
	for _, m := range ctxLog.MessageLog {
		fmt.Fprintf(w, fmt.Sprintf("[%s] %s: %s<br>", m.Timestamp.Format("15:04:05"), m.User, m.Content))
	}
	fmt.Fprintf(w, HtmlLogEnd)
}
