package handlers

import (
	"strings"

	"github.com/Fabriciope/cli_chat/client/clientapp/cui"
	"github.com/Fabriciope/cli_chat/shared"
)

func (handler *Handler) NewClientResponseHandler(response shared.Response) {
	chatLine := cui.ChatLine{
		Info:      "[insert time]",
		InfoColor: shared.Green,
		Text:      strings.Trim(response.Payload.(string), " "),
	}
	handler.CUI().DrawNewLineInChat(&chatLine)
}
