package responsehandler

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

// TODO: não está funcionando perfeitamente
func (handler *ResponseHandler) GetUsers(response dto.Response) {
	if response.Err {
		handler.cui.PrintLine(&cui.Line{
			Info:      "ERROR FROM SERVER:",
			Text:      response.Payload.(string),
			TextColor: escapecode.Red,
		})
		return
	}

	var users []map[string]string
	err := json.Unmarshal([]byte(response.Payload.(string)), &users)
	if err != nil {
		handler.cui.PrintLine(&cui.Line{
			Text:      response.Payload.(string),
			TextColor: escapecode.Yellow,
		})
		return
	}

	var usersName string
	for _, user := range users {
		usersName += fmt.Sprintf(" %s%s%s,", user["color"], user["name"], escapecode.Reset)
	}

	usersName = fmt.Sprintf("[%s ]", strings.TrimSuffix(usersName, ","))
	handler.cui.PrintLine(&cui.Line{
		Info:      "users:",
		InfoColor: escapecode.Blue,
		Text:      usersName,
	})
}
