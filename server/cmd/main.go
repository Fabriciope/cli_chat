package main

import "github.com/Fabriciope/cli_chat/server/serverapp"

func main() {
	server := serverapp.NewServer()
	server.InitServer()
}
