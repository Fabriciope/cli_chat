package main

import (
	"log"

	"github.com/Fabriciope/cli_chat/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Panic(err)
	}

	server.InitServer()
}
