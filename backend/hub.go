package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Hub to take care of all channels and sockets
type Hub struct {
	channels map[string]*Room
	upgrader websocket.Upgrader
}

// constructor for hub
func makeHub() *Hub {
	hub := new(Hub)
	hub.channels = make(map[string]*Room)
	hub.upgrader = websocket.Upgrader{}

	return hub
}

func (h *Hub) checkChannel(name string) *Room {
	if _, ok := h.channels[name]; !ok {
		h.channels[name] = makeRoom(name)
	}
	return h.channels[name]
}

func (h *Hub) handleWebSockets(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	channelName := path["channel"]

	h.upgrader.CheckOrigin = func(*http.Request) bool { return true } // allow requests from wherever
	ws, err := h.upgrader.Upgrade(w, r, nil)                          // upgrade http request to web socket
	if err != nil {
		log.Fatal("Upgrade: ", err)
	}

	defer ws.Close()

	channel := h.checkChannel(channelName)
	id := h.channels[channelName].joinRoom()

	for {
		var msg Message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error ReadJSON: %v", err)
			delete(h.channels[channelName].clients, id)
			break
		}

		channel.clients[id].channel <- msg
	}
}
