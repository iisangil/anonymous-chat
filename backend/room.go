package main

import (
	"log"
)

// Room for messaging
type Room struct {
	name    string
	clients map[int]*Client
	index   int
}

// constructor for channels
func makeRoom(name string) *Room {
	room := new(Channel)
	room.name = name
	room.clients = make(map[int]*Client)
	room.index = 0

	return room
}

func (c *Room) joinRoom() int {
	c.index++
	client := makeClient(c.index)
	c.clients[c.index] = client
	return c.index
}

func (c *Room) handleMessages(id int) {
	for {
		msg := <-c.clients[id].channel

		for client := range c.clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Print("error write JSON: %v", err)
				client.Close()
				delete(c.clients, client)
			}
		}
	}
}
