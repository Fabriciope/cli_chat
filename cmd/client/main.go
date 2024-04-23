package main

import (
	"log"
	"os"

	"github.com/Fabriciope/cli_chat/client"
	"github.com/Fabriciope/cli_chat/client/cui"
)

const (
	remoteIp   = "cli_chat-server"
	remotePort = 5000
)

func main() {
	userInterface, err := cui.NewCUI()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	user, err := client.NewUser(remoteIp, remotePort, userInterface)
	if err != nil {
		userInterface.PrintLineForInternalError(err.Error())
		return
	}

	go userInterface.InitConsoleUserInterface()
	user.Login()
	user.InitChat()
}
