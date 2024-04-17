package responsehandler

import (
	"encoding/json"
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
		handler.cui.PrintLine(
			cui.MakeLine(&cui.Line{
				Info:      "login status:",
				InfoColor: escapecode.BrightYellow,
				Text:      response.Payload.(string),
				TextColor: escapecode.Yellow,
			}))

		return
	}

	*handler.userLoggedIn = true
	handler.cui.RenderChatInterface()
	handler.cui.PrintLine(
		cui.MakeLine(&cui.Line{
			Info:      "login status:",
			InfoColor: escapecode.BrightGreen,
			Text:      response.Payload.(string),
			TextColor: escapecode.Green, // TODO: testar sem co TextColor
		}))
}

func (handler *ResponseHandler) NewMessageReceived(response dto.Response) {
	var textMessage dto.TextMessage
	json.Unmarshal([]byte(response.Payload.(string)), &textMessage)
	handler.cui.PrintLine(
		cui.MakeLine(&cui.Line{
			Info:      escapecode.TextToBold(textMessage.Username + ":"),
			InfoColor: textMessage.UserColor,
			Text:      textMessage.Message,
		}))
}

func (handler *ResponseHandler) NewClient(response dto.Response) {
	handler.cui.PrintLine(
		cui.MakeLine(&cui.Line{
			InfoColor: escapecode.BrightGreen,
			Text:      strings.Trim(response.Payload.(string), " "),
			TextColor: escapecode.Green,
		}))
}

func (handler *ResponseHandler) SendMessageInChatResponse(response dto.Response) {
	if response.Err {
		handler.cui.PrintLine(
			cui.MakeLine(&cui.Line{
				Info:      "ERROR FROM SERVER:",
				InfoColor: escapecode.Red,
				Text:      response.Payload.(string),
				TextColor: escapecode.Yellow,
			}))
	}
}
