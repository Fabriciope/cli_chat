package responsehandler

import (
	"github.com/Fabriciope/cli_chat/client/cui"
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
