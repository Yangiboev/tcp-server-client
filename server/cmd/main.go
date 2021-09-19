package main

import (
	"github.com/Yangiboev/tcp-server-client/server"
)

func main() {

	s := server.NewServer()
	s.Listen(":5000")
	// start the server
	go s.Start()
	server.StartUIServer(s)
}
