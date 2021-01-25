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

// constructor for hub
func makeHub() *Hub {
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
	channelName := path["channel"]

	h.upgrader.CheckOrigin = func(*http.Request) bool { return true } // allow requests from wherever
	ws, err := h.upgrader.Upgrade(w, r, nil)                          // upgrade http request to web socket
	if err != nil {
		log.Fatal("Upgrade: ", err)
	}

	defer ws.Close()

	h.checkChannel(channelName)
	h.channels[channelName].JoinChannel(ws)

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
