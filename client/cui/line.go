package cui

import (
	"fmt"
	"strings"
	"time"

	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

type Line struct {
	Info      string
	InfoColor escapecode.ColorCode
	Text      string
	TextColor escapecode.ColorCode
}

// TODO: nao retornar line, somente fazer referencia ao valor
func addDataToLine(line *Line) *Line {
	timeStr := time.Now().Format(time.TimeOnly)
	line.Info = strings.Trim(fmt.Sprintf("[%s] %s", timeStr, line.Info), " ")

	// TODO: se tiver so o texto, colocar a cor do info a mesma do textcolor

	if line.TextColor == "" {
		line.TextColor = escapecode.White
	}

	return line
}
