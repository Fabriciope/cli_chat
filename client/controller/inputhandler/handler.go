package inputhandler

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/client/sender"
)

var AvailableCommands = [...]string{
    ":h", // NOTE: help
    ":q", // NOTE: quit
    ":users", // NOTE: number of users
	":numberOfUsers",
}

const (
	H int = iota
	Q
	Users
	NumberOfUsers
)

type InputHandler struct {
	userLoggedIn *bool
	cui          cui.CUIInterface
	sender       sender.SenderInterface
}

func NewInputHandler(conn *net.TCPConn, cui cui.CUIInterface, loggedIn *bool) *InputHandler {
	return &InputHandler{
		userLoggedIn: loggedIn,
		cui:          cui,
		sender:       sender.NewRequestSender(conn),
	}
}
