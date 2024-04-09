package cui

import (
	"github.com/Fabriciope/cli_chat/pkg/escapecode"
)

type ChatLine struct {
	Info      string
	InfoColor escapecode.ColorCode
	Text      string // TODO: adicionar uma cor para o Text
}
