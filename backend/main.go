package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Message struct to hold message information
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Channel  string `json:"channel"`
}

// var clients = make(map[*websocket.Conn]bool) // connected clients
// var broadcast = make(chan Message)           // channel to broadcast messages
// var upgrader = websocket.Upgrader{}          // upgrader for websockets

func main() {
	hub := makeHub()
	r := mux.NewRouter()
	r.HandleFunc("/ws/{room}", hub.handleWebSockets)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// func main() {
// 	http.HandleFunc("/ws", handleWebSockets) // websocket initiation route
// 	go handleMessages()                      // goroutine to handle sending out messages
// 	err := http.ListenAndServe(":8000", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

/**
func handleWebSockets(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(*http.Request) bool { return true } // allow requests from wherever
	ws, err := upgrader.Upgrade(w, r, nil)                          // upgrade http request to web socket
	if err != nil {
		log.Fatal("Upgrade: ", err)
	}

	defer ws.Close() // make sure to close websocket once functions finishes

	clients[ws] = true // add new websocket to global map

	// loop that accepts messages and broadcasts to channel
	for {
		var msg Message

		err := ws.ReadJSON(&msg) // read in message and parse as json
		if err != nil {
			log.Printf("error ReadJSON: %v", err)
			delete(clients, ws) // remove websocket from global map
			break
		}
		fmt.Println("got a message")
		fmt.Printf("%v", msg)

		broadcast <- msg // send received message to broadcast channel
	}
}

func handleMessages() {
	for {
		msg := <-broadcast // get message from channel
		fmt.Println("message from broadcast channel")
		fmt.Printf("%v", msg)

		// send message to connected web sockets
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error WriteJSON: %v", err)
				client.Close()          // close web socket connection
				delete(clients, client) // remove from global map
			}
		}
	}
}
**/
