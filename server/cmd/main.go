package main

import (
	"log"

	"github.com/Fabriciope/cli_chat/server/serverapp"
)

func main() {
	server, err := serverapp.NewServer()
	if err != nil {
		log.Panicln(err)
		return
	}
	
	server.InitServer()
}
