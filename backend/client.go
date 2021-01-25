package main

import "github.com/gorilla/websocket"

// Client sends and receives messages
type Client struct {
	id      int
	ws      *websocket.Conn
	channel chan Message
}

func makeClient(id int) *Client {
	client := new(Client)
	client.id = id
	client.channel = make(chan Message)
}

func (c *Client) sendMessage(msg Message) {
	c.channel <- msg
}
