package main

import (
	"log"
	"os"

	"github.com/Fabriciope/cli_chat/client"
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

func main() {
	userInterface, err := cui.NewCUI()
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	user, err := client.NewUser(userInterface)
	if err != nil {
		userInterface.PrintLineAndExit(1, cui.Line{
			Info:      "error creating client:",
			InfoColor: escapecode.Red,
			Text:      err.Error(),
			TextColor: escapecode.Yellow,
		})

		return
	}

	user.InitChat() // TODO: colorcar instrucao dentro de um log para verificar os errors
}
