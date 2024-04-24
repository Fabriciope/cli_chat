package main

import (
	"flag"
	"log"

	"github.com/Fabriciope/cli_chat/server"
)

func main() {
	var (
		ip   = "0.0.0.0"
		port = 5000
	)

	flag.StringVar(&ip, "ip", ip, "IP address to listen on")
	flag.IntVar(&port, "port", port, "Port to listen on")

	flag.Parse()

	server, err := server.NewTCPServer(ip, port)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	server.InitServer()
}
