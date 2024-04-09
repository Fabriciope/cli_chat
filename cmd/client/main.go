package main

import (
	"log"

	"github.com/Fabriciope/cli_chat/client"
)

func main() {
	user, err := client.NewUser()
	if err != nil {
		log.Panic(err)
	}

	user.InitChat()
}
