package responsehandler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

type ResponseHandler struct {
	userLoggedIn *bool
	cui          cui.CUIInterface
}

func NewResponseHandler(cui cui.CUIInterface, loggedIn *bool) *ResponseHandler {
	return &ResponseHandler{
		userLoggedIn: loggedIn,
		cui:          cui,
	}
}

func (handler *ResponseHandler) LoginResponse(response dto.Response) {
	if response.Err {
		handler.cui.DrawLoginError(response.Payload.(string))
		return
	}

	*handler.userLoggedIn = true
	handler.cui.SetLoggedAs(true)
	handler.cui.DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time] login status:",
		InfoColor: escapecode.BrightGreen,
		Text:      response.Payload.(string),
	})
	// TODO: tirar o gerecianmento do troca de interface do cui para a interface
}

func (handler *ResponseHandler) NewMessageReceived(response dto.Response) {
	var textMessage dto.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	handler.cui.DrawNewLineInChat(&cui.ChatLine{
		// TODO: colocar o username em bold
		Info:      fmt.Sprintf("[insert time] %s:", textMessage.Username),
		InfoColor: textMessage.UserColor,
		Text:      textMessage.Message,
	})
}

func (handler *ResponseHandler) NewClient(response dto.Response) {
	handler.cui.DrawNewLineInChat(&cui.ChatLine{
		Info:      "[insert time]",
		InfoColor: escapecode.Green,
		Text:      strings.Trim(response.Payload.(string), " "),
	})
}

func (handler *ResponseHandler) SendMessageInChatResponse(response dto.Response) {
	if response.Err {
		handler.cui.DrawNewLineInChat(&cui.ChatLine{
			Info:      "[insert time] ERROR FROM SERVER:",
			InfoColor: escapecode.Red,
			Text:      response.Payload.(string),
		})
	}
}
