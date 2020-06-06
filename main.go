package main

import (
	"github.com/dharmasastra/realtime-chat-go/app/controllers"
	"github.com/dharmasastra/realtime-chat-go/config"
)

func main()  {
	e := config.NewRouter()

	// Create a simple file server
	e.Static("/", "./public")

	// Start listening for incoming chat messages
	go controllers.HandleMessages()

	// Start the server on localhost port 8080 and log any errors
	e.Logger.Fatal(e.Start(":8080"))
}
