package inputhandler

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/client/sender"
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

// TODO: adicionar um input para e comando :availableCommands
