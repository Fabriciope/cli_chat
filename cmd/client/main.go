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
		userInterface.DrawLineAndExit(1, cui.ChatLine{
			Info:      "[insert time]",
			InfoColor: escapecode.Red,
			Text:      err.Error(),
		})
	}

	user.InitChat()
}
