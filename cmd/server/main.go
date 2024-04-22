package main

import (
	"log"

	"github.com/Fabriciope/cli_chat/server"
)

const (
	ip   = "localhost"
	port = 5000
)

func main() {
	server, err := server.NewTCPServer(ip, port)
	if err != nil {
		log.Panic(err)
	}

	server.InitServer()
}
