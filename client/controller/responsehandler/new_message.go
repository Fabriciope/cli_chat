package responsehandler

import (
	"encoding/json"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) NewMessageReceived(response dto.Response) {
	var textMessage dto.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	handler.cui.PrintLine(&cui.Line{
		Info:      escapecode.TextToBold(textMessage.Username + ":"),
		InfoColor: textMessage.UserColor,
		Text:      textMessage.Message,
	})
}
