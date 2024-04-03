package handlers

import (
	"errors"
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)


// TODO: resolver problema quando a escrita do usuario ultrapassa o consoleWidth
// TODO: definir uma cor para cada client
func (handler *Handler) SendMessageInChat(message string) error {
	if !handler.user.LoggedIn() {
		return errors.New("you must be logged in to send messages in chat")
	}

	// TODO: verificar se a mensagem esta vazia
	err := handler.sender.SendRequest(shared.Request{
		Name:    shared.SendMessageActionName,
		Payload: strings.Trim(message, " "),
	})
	if err != nil {
		return err
	}

	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me: ",
		InfoColor: cui.Yellow,
		Text:      message,
	})
	return nil
}

func (handler *Handler) SendMessageInChatResponse(response shared.Response) {
	if response.Err {
		handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] ",
			InfoColor: cui.Red,
			Text:      response.Payload.(string),
		})
	}
}
