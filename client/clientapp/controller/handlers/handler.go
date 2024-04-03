package handlers

import (
	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/client/clientapp/interfaces"
	"github.com/Fabriciope/cli_chat/client/clientapp/sender"
)

type Handler struct {
	user   interfaces.Client
	sender interfaces.Sender
}

func NewHandler(user interfaces.Client) *Handler {
	return &Handler{
		user:   user,
		sender: sender.NewRequestSender(user.Conn()),
	}
}

func (handler *Handler) CommandHandler(handlerName string) interfaces.CommandHandler {
	return nil
}

func (handler *Handler) ResponseHandler(string) interfaces.ResponseHandler {
	return nil
}

func (handler *Handler) CUI() *cui.CUI {
	return handler.user.CUI()
}
