package main

import (
	"flag"
	"log"

	"github.com/Yangiboev/tcp-server-client/client"
	"github.com/Yangiboev/tcp-server-client/ui"
)

func main() {
	address := flag.String("server", "localhost:5000", "Which server to connect to")
	flag.Parse()
	client := client.NewClient()
	err := client.Dial(*address)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// start the client to listen for incoming message
	go client.Start()
	ui.StartUi(client)
}
