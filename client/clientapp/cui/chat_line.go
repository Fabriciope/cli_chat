package cui

import "github.com/Fabriciope/cli_chat/shared"

type ChatLine struct {
	Info      string
	InfoColor shared.ColorCode
    Text      string // TODO: adicionar uma cor para o Text
}
