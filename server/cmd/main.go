package main

import (
	"github.com/Yangiboev/chatserver/server"
)

func main() {
	var s server.ChatServer
	s = server.NewServer()
	s.Listen(":5000")

	// start the server
	s.Start()
}
