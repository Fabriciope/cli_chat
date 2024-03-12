package main

import "github.com/Fabriciope/cli_chat/server/serverapp"

func main() {
	var server = serverapp.NewServer()
	server.InitServer()
}