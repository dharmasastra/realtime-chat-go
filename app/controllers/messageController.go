package controllers

import (
	"github.com/dharmasastra/realtime-chat-go/app/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // connection clients
var broadcast = make(chan models.Message)	// broadcast chanel
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}	// configure the upgrader

func HandleConnections(c echo.Context) error {
	// Upgrader initial Get request to a websocket
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
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
			//log.Printf("error: %v", err)
			c.Logger().Error(err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast chanel
		broadcast <- msg
	}
	return nil
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