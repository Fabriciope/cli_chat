package inputhandler

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/client/sender"
)

var AvailableCommands = [...]string{
	":h",
	":help",
	":q",
	":quit",
	":logout",
	":users",
	":numberOfUsers",
	":availableCommands",
}

const (
	H int = iota
	Help
	Q
	Quit
	Logout
	Users
	NumberOfUsers
	Commands
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
