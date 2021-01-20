package main

import "github.com/gorilla/websocket"

// Channel for messaging
type Channel struct {
	name    string
	clients map[*websocket.Conn]bool
	channel chan message
}

// MakeChannel constructor for channels
func MakeChannel(name string) *Channel {
	channel := new(Channel)
	channel.name = name
	channel.clients = make(map[*websocket.Conn]bool)
	channel.channel = make(chan message)

	return channel
}
