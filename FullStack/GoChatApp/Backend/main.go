package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrader(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setuproutes() {
	// Handles incoming requests and written responses and prints it out
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Chat server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Chat server v0.01")
	setuproutes()
	http.ListenAndServe(":8080", nil)
}
