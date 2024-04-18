package inputhandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *InputHandler) SendMessageInChat(message string) {
	if !*handler.userLoggedIn {
		handler.cui.PrintLine(&cui.Line{
			Info:      "warning:",
			InfoColor: escapecode.BrightYellow,
			Text:      "you must be logged in to send messages in chat",
			TextColor: escapecode.Yellow,
		})

		return
	}

	request := dto.Request{Name: dto.SendMessageActionName, Payload: message}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.cui.PrintLineForInternalError(err.Error())
		return
	}

	handler.cui.PrintLine(&cui.Line{
		Info:      escapecode.TextToBold("me:"),
		InfoColor: escapecode.DefaultColor,
		Text:      message,
	})
}
