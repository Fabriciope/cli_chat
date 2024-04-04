package handlers

import (
	"encoding/json"
	"fmt"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)

func (handler *Handler) NewMessageReceivedHandler(response shared.Response) {
	var textMessage shared.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	chatLine := cui.ChatLine{
		Info:      fmt.Sprintf("[insert time] From %s:", textMessage.Username),
		InfoColor: shared.BrightYellow,
		Text:      textMessage.Message,
	}
	handler.CUI().DrawNewLineInChat(&chatLine)
}
