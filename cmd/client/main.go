package main

import (
	"log"
	"os"

	"github.com/Fabriciope/cli_chat/client"
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
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
		userInterface.PrintLineAndExit(1, cui.Line{
			Info:      "error creating client:",
			Text:      err.Error(),
			TextColor: escapecode.Red,
		})

		return
	}

	go userInterface.InitConsoleUserInterface()
	user.Login()
	user.InitChat()
}
