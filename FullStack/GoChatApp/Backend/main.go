package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//Checks origin of messages
//Allows us to make requests to React
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	CheckOrigin: func(r *http.Request) bool { return true },
}

// read messages sent by client and prints
func readmessage(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

	}
}

//Defining our endpoint
//Also upgrades connection to websocket
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	readmessage(ws)
}

func setuproutes() {
	// Handles incoming requests and written responses and prints it out using function earlier
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Chat server")
	})

	http.HandleFunc("/ws", serveWs)
}

//opens live server

func main() {
	fmt.Println("Chat server v0.01")
	setuproutes()
	http.ListenAndServe(":8080", nil)
}
