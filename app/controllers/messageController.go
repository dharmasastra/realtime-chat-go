package controllers

import (
	"github.com/dharmasastra/realtime-chat-go/app/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
)

var clients = make(map[*websocket.Conn]bool) // connection clients
var broadcast = make(chan models.Message)	// broadcast chanel
var upgrader = websocket.Upgrader{}	// configure the upgrader

func HandleConnections(c echo.Context)  {
	// Upgrader initial Get request to a websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when function returns
	defer ws.Close()

	// Rgister our new client
	clients[ws] = true

	for {
		var msg models.Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast chanel
		broadcast <- msg
	}
}

func HandleMessages()  {
	for {
		// Grab the next message from the broadcast chanel
		msg := <- broadcast
		// Send it out to every client that is currently connected
		for client := range clients{
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}