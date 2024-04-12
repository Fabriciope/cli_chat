package handler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared"
)

func (handler *Handler) SendMessageInChat(message string) {
	if !handler.user.LoggedIn() {
		handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] Me:",
			InfoColor: escapecode.Yellow,
			Text:      "you must be logged in to send messages in chat",
		})
		return
	}

	request := shared.Request{
		Name:    shared.SendMessageActionName,
		Payload: message,
	}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.CUI().DrawNewLineForInternalError()
	}

	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me:",
		InfoColor: escapecode.Yellow,
		Text:      message,
	})
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
