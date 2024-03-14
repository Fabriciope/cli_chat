package main

import "github.com/Fabriciope/cli_chat/client/clientapp"

func main() {
	client := clientapp.NewClient()
	client.InitChat()
}