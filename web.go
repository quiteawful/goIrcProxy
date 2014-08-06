package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func startWebServer() {
	http.HandleFunc("/", HttpIndex)
	http.HandleFunc("/log", HttpLog)
	http.ListenAndServe(":8080", nil)
}

func checkAuth(w http.ResponseWriter, r *http.Request) bool {
	s := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(s) != 2 {
		return false
	}

	b, err := base64.StdEncoding.DecodeString(s[1])
	if err != nil {
		return false
	}

	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return false
	}

	return pair[0] == "doclol" && pair[1] == "g0IrcProxy"
}

// Serves the main index page, pretty much only the messagelog
func HttpIndex(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) { // auth
		content := r.FormValue("content")

		if content != "" { // is POST
			ctxLog.AddWebLog("doclol", content)
		}

		fmt.Fprintf(w, HtmlMain)
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="schleich dich"`)
	//w.Write([]byte("401"))
	w.Write([]byte("401 Unauthorized\n"))
}

func HttpLog(w http.ResponseWriter, r *http.Request) {
	if checkAuth(w, r) { // auth
		for _, m := range ctxLog.MessageLog {
			fmt.Fprintf(w, fmt.Sprintf("[%s] %s: %s<br>", m.Timestamp.Format("15:04:05"), m.User, m.Content))
		}
		return
	}

	w.Header().Set("WWW-Authenticate", `Basic realm="schleich dich"`)
	//w.Write([]byte("401"))
	w.Write([]byte("401 Unauthorized\n"))
}
