package responsehandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) Logout(response dto.Response) {
	if response.Err {
		handler.cui.PrintLine(&cui.Line{
			Info:      "logout err:",
			InfoColor: escapecode.Yellow,
			Text:      response.Payload.(string),
			TextColor: escapecode.Yellow,
		})
		return
	}

	*handler.userLoggedIn = false
	handler.cui.PrintLineAndExit(0, cui.Line{
		InfoColor: escapecode.Blue,
		Text:      "you have been disconnected",
		TextColor: escapecode.Blue,
	})
}
