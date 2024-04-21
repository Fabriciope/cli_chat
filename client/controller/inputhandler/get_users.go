package inputhandler

import "github.com/Fabriciope/cli_chat/pkg/shared/dto"

func (handler *InputHandler) GetUsers(arg string) {
	err := handler.sender.SendRequest(dto.Request{Name: dto.GetUsersActionName})
	if err != nil {
		handler.cui.PrintLineForInternalError(err.Error())
		return
	}
}
