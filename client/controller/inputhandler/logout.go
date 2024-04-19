package inputhandler

import (
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *InputHandler) Logout(arg string) {
	err := handler.sender.SendRequest(dto.Request{Name: dto.LogoutActionName})
	if err != nil {
		handler.cui.PrintLineForInternalError(err.Error())
		return
	}
}
