package inputhandler

import (
	"fmt"
	"strings"

	"github.com/Fabriciope/cli_chat/client/cui"
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

func (handler *InputHandler) GetAvailableCommands(arg string) {
	var commands string
	for _, command := range AvailableCommands {
		commands += fmt.Sprintf(" %s,", command)
	}

	handler.cui.PrintLine(&cui.Line{
		Info:      "Available commmands:",
		Text:      fmt.Sprintf("[%s ]", strings.TrimSuffix(commands, ",")),
		TextColor: escapecode.Blue,
	})
}
