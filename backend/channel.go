package main

import "github.com/gorilla/websocket"

// Channel for messaging
type Channel struct {
	name    string
	clients map[*websocket.Conn]bool
	channel chan message
}

// constructor for channels
func makeChannel(name string) *Channel {
	channel := new(Channel)
	channel.name = name
	channel.clients = make(map[*websocket.Conn]bool)
	channel.channel = make(chan message)

	return channel
}

func (c *Channel) joinChannel(ws *websocket.Conn) {
	c.clients[ws] = true // add new websocket to room
}

func (c *Channel) HandleMessages() {
	for {
		msg := <-c.channel

		for client := range clients {

		}
	}
}
