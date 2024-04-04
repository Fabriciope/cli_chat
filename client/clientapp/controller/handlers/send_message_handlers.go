package handlers

import (
	"errors"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)

// TODO: resolver problema quando escreve e acumula o texto
func (handler *Handler) SendMessageInChat(message string) error {
	if !handler.user.LoggedIn() {
		return errors.New("you must be logged in to send messages in chat")
	}

	// TODO: verificar se a mensagem esta vazia
	err := handler.sender.SendRequest(shared.Request{
		Name:    shared.SendMessageActionName,
		Payload: message,
	})
	if err != nil {
		return err
	}

	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me:",
		InfoColor: shared.Yellow,
		Text:      message,
	})
	return nil
}

func (handler *Handler) SendMessageInChatResponse(response shared.Response) {
	if response.Err {
		handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] ERROR FROM SERVER:",
			InfoColor: shared.Red,
			Text:      response.Payload.(string),
		})
	}
}
