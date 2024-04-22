package responsehandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) SendMessageInChat(response dto.Response) {
	if response.Err {
		handler.cui.PrintLine(&cui.Line{
			Info:      "ERROR FROM SERVER:",
			Text:      response.Payload.(string),
			TextColor: escapecode.Red,
		})
	}
}
