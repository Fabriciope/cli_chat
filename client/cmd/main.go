package main

import (
	"log"

	"github.com/Fabriciope/cli_chat/client/clientapp"
)

func main() {
	client, err := clientapp.NewUser()
	if err != nil {
		log.Panicln(err)
		return
	}

	client.InitChat()
}
