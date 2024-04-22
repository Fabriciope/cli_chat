package inputhandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *InputHandler) Login(username string) {
	if *handler.userLoggedIn {
		handler.cui.PrintLine(&cui.Line{
			Info:      "invalid operation:",
			Text:      "user is already logged in",
			TextColor: escapecode.Red,
		})
		return
	}

	request := dto.Request{Name: dto.LoginActionName, Payload: username}
	err := handler.sender.SendRequest(request)
	if err != nil {
		handler.cui.PrintLineForInternalError(err.Error())
		return
	}
}
