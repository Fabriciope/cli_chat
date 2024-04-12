package handler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
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

	request := dto.Request{
		Name:    dto.SendMessageActionName,
		Payload: message,
	}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.CUI().DrawNewLineForInternalError(err.Error())
	}

	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] Me:",
		InfoColor: escapecode.Yellow,
		Text:      message,
	})
}

func (handler *Handler) SendMessageInChatResponse(response dto.Response) {
	if response.Err {
		handler.CUI().DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] ERROR FROM SERVER:",
			InfoColor: escapecode.Red,
			Text:      response.Payload.(string),
		})
	}
}
