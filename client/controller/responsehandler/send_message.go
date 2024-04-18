package responsehandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) SendMessageInChatResponse(response dto.Response) {
	if response.Err {
		handler.cui.PrintLine(&cui.Line{
			Info:      "ERROR FROM SERVER:",
			InfoColor: escapecode.Red,
			Text:      response.Payload.(string),
			TextColor: escapecode.Yellow,
		})
	}
}
