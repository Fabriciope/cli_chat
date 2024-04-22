package responsehandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) UserDisconnected(response dto.Response) {
	handler.cui.PrintLine(&cui.Line{
		Text:      response.Payload.(string),
		TextColor: escapecode.Blue,
	})
}
