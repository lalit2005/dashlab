package main

import (
	"dashlab/client"
	"dashlab/server"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("Argument missing")
	}
	if os.Args[1] == "server" {
		server.StartServer()
	} else if os.Args[1] == "client" {
		fmt.Println("client started")
		client.StartClient()
	} else {
		fmt.Errorf("Please provide a valid argument: client or server")
	}
}
