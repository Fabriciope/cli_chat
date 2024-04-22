package responsehandler

import (
	"strings"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) NewClient(response dto.Response) {
	handler.cui.PrintLine(&cui.Line{
		Text:      strings.Trim(response.Payload.(string), " "),
		TextColor: escapecode.Green,
	})
}
