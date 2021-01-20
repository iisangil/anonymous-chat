package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

// Hub to take care of all channels and sockets
type Hub struct {
	channels map[string]*Channel
	upgrader websocket.Upgrader
}

// MakeHub constructor for hub
func MakeHub() *Hub {
	hub := new(Hub)
	hub.channels = make(map[string]*Channel)
	hub.upgrader = websocket.Upgrader{}

	return hub
}

func (h *Hub) checkChannel(name string) {
	if _, ok := h.channels[name]; !ok {
		h.channels[name] = MakeChannel(name)
	}
}

func (h *Hub) handleWebSockets(w http.ResponseWriter, r *http.Request) {
	path := mux.Vars(r)
	var channelName string
	if channel, ok := path["channel"]; !ok {
		channelName = "global"
	} else {
		channelName = channel
	}

	h.checkChannel(channelName)

	h.upgrader.CheckOrigin = func(*http.Request) bool { return true } // allow requests from wherever
	ws, err := h.upgrader.Upgrade(w, r, nil)                          // upgrade http request to web socket
	if err != nil {
		log.Fatal("Upgrade: ", err)
	}

	defer ws.Close()

	h.channels[channelName].clients[ws] = true // add new websocket to roo

	for {
		var msg message

		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error ReadJSON: %v", err)
			delete(h.channels[channelName].clients, ws)
			break
		}

		h.channels[channelName].channel <- msg
	}
}

func (h *Hub) handleMessages() {
	for {

	}
}
