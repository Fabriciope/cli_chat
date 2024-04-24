package main

import (
	"flag"
	"log"
	"os"

	"github.com/Fabriciope/cli_chat/client"
	"github.com/Fabriciope/cli_chat/client/cui"
)

func main() {
	var (
		remoteIp   = ""
		remotePort = 5000
	)

	flag.StringVar(&remoteIp, "ip", remoteIp, "Remote server IP")
	flag.IntVar(&remotePort, "port", remotePort, "Remote server port")

	flag.Parse()

	if remoteIp == "" {
		log.Print("Remote IP is required")
		os.Exit(1)
	}

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
