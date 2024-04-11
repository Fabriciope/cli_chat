package handler

import (
	"errors"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

func (handler *Handler) SendMessageInChat(message string) error {
	if !handler.user.LoggedIn() {
		return errors.New("you must be logged in to send messages in chat")
	}

	err := handler.sender.SendRequest(shared.Request{
		Name:    shared.SendMessageActionName,
		Payload: message,
	})
	if err != nil {
		return err
	}

	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me:",
		InfoColor: escapecode.Yellow,
		Text:      message,
	})
	return nil
}

func (handler *Handler) SendMessageInChatResponse(response shared.Response) {
	if response.Err {
		handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] ERROR FROM SERVER:",
			InfoColor: escapecode.Red,
			Text:      response.Payload.(string),
		})
	}
}
