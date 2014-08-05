package main

import (
	"fmt"
)

var FormTemplate string = "<form action='/write' method='post'><input type='text' name='user'><input type='text' name='content'><input type='submit'></form>"

func main() {
	fmt.Println("Start Proxy")
	//MessageLog = append(MessageLog, &Message{Timestamp: time.Now(), User: "marduk", Content: "hi"})
	go startWebServer()

	ch := make(chan bool)
	<-ch
}
