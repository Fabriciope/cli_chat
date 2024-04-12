package handler

import (
	"strings"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *Handler) NewClientResponseHandler(response dto.Response) {
	chatLine := cui.ChatLine{
		Info:      "[insert time]",
		InfoColor: escapecode.Green,
		Text:      strings.Trim(response.Payload.(string), " "),
	}
	handler.CUI().DrawNewLineInChat(&chatLine)
}
