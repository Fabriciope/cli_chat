package handler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/client/interfaces"
	"github.com/Fabriciope/cli_chat/client/sender"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

type CommandHandler func() error
type CommandsHandlersMap map[string]CommandHandler
type ResponseHandler func(shared.Response)
type ResponsesHandlersMap map[string]ResponseHandler

type Handler struct {
	user   interfaces.Client
	sender interfaces.Sender
}

// TODO: dividir os handlers entre input e commands
func NewHandler(user interfaces.Client) *Handler {
	return &Handler{
		user:   user,
		sender: sender.NewRequestSender(user.Conn()),
	}
}

func (handler *Handler) CommandHandler(handlerName string) CommandHandler {
	return nil
}

func (handler *Handler) ResponseHandler(string) ResponseHandler {
	return nil
}

func (handler *Handler) CUI() *cui.CUI {
	return handler.user.CUI()
}
