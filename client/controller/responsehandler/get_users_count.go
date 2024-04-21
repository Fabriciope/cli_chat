package responsehandler

import (
	"fmt"
	"strconv"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
	"github.com/Fabriciope/cli_chat/pkg/shared/dto"
)

func (handler *ResponseHandler) GetUsersCount(response dto.Response) {
	numberOfUsers, _ := strconv.Atoi(response.Payload.(string))
	if numberOfUsers == 0 {
		handler.cui.PrintLine(&cui.Line{
			InfoColor: escapecode.Blue,
			Text:      "you are the only user in this room",
			TextColor: escapecode.Blue,
		})
		return
	}

	handler.cui.PrintLine(&cui.Line{
		InfoColor: escapecode.Blue,
		Text:      fmt.Sprintf("the number of users in the room is %d", numberOfUsers),
		TextColor: escapecode.Blue,
	})
}
