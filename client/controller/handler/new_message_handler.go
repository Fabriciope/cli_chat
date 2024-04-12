package handler

import (
	"encoding/json"
	"fmt"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *Handler) NewMessageReceivedHandler(response dto.Response) {
	var textMessage dto.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	handler.CUI().DrawNewLineInChat(&cui.ChatLine{
		// TODO: colocar o username em bold
		Info:      fmt.Sprintf("[insert time] %s:", textMessage.Username),
		InfoColor: textMessage.UserColor,
		Text:      textMessage.Message,
	})
}
