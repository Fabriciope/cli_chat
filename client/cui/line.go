package cui

import (
	"strings"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

type Line struct {
	Info      string
	InfoColor escapecode.ColorCode
	Text      string
	TextColor escapecode.ColorCode
}

// TODO: testar na aplicacao
func MakeLine(line *Line) *Line {
	// TODO: implementar, pegar o tempo atual quando a struct Line for declarada
	line.Info = strings.Trim("[insert time] "+line.Info, " ")

	if line.TextColor == "" {
		line.TextColor = escapecode.White
	}

	return line
}
