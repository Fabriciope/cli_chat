package inputhandler

import (
	"net"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/client/sender"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type InputHandler struct {
	userLoggedIn *bool
	cui          cui.CUIInterface
	sender       sender.SenderInterface
}

// TODO: colocar cada handler em um arquivo

func NewInputHandler(conn *net.TCPConn, cui cui.CUIInterface, loggedIn *bool) *InputHandler {
	return &InputHandler{
		userLoggedIn: loggedIn,
		cui:          cui,
		sender:       sender.NewRequestSender(conn),
	}
}

func (handler *InputHandler) Login(username string) {
	if username == "" {
		handler.cui.DrawLoginError("write something")
		return
	}

	// TODO: verificar se o usuario ja esta logado
	request := dto.Request{Name: dto.LoginActionName, Payload: username}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.cui.DrawNewLineForInternalError(err.Error())
		return
	}
}

func (handler *InputHandler) SendMessageInChat(message string) {
	if !*handler.userLoggedIn {
		handler.cui.DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] Me:",
			InfoColor: escapecode.Yellow,
			Text:      "you must be logged in to send messages in chat",
		})
		return
	}

	request := dto.Request{
		Name:    dto.SendMessageActionName,
		Payload: message,
	}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.cui.DrawNewLineForInternalError(err.Error())
		return
	}

	handler.cui.DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me:",
		InfoColor: escapecode.Yellow,
		Text:      message,
	})
}
